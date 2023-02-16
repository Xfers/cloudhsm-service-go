package api

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/Xfers/cloudhsm-service-go/api/controllers"
	"github.com/Xfers/cloudhsm-service-go/crypto"
	_ "github.com/Xfers/cloudhsm-service-go/docs"
	"github.com/Xfers/go-openssl"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	BASE_ROUTE = "api"
	PORT       = "8000"
)

// @title CloudHSM service API
// @version 1.0
// @description Standalone cloudhsm-service

// @host localhost:8000
// @BasePath /api/
// @schemes http
func RunSignerServer(flags map[string]interface{}) {
	err := prepareSignerServer(flags)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	setSignerRoutes(r)

	// Run server
	err = r.Run(":" + PORT)
	if err != nil {
		panic(err)
	}
}

func prepareSignerServer(flags map[string]interface{}) error {

	// Get keys from flags
	flagKeys := flags["keys"].(map[string]string)
	keys, err := getKeys(flagKeys, "private")
	if err != nil {
		return err
	}

	controllers.NewBaseController(&keys)
	return nil
}

func setSignerRoutes(r *gin.Engine) {

	setBaseRoutes(r)

	// Digest endpoint with BASE_ROUTE
	endPoint := BASE_ROUTE + "/digest"
	r.POST(endPoint, func(c *gin.Context) {
		controllers.DigestController(c)
	})
	// Sign endpoint
	endPoint = BASE_ROUTE + "/sign/:keyName"
	r.POST(endPoint, func(c *gin.Context) {
		controllers.SignController(c)
	})

	// Pure Sign endpoint
	endPoint = BASE_ROUTE + "/pure-sign/:keyName"
	r.POST(endPoint, func(c *gin.Context) {
		controllers.PureSignController(c)
	})
}

func RunVerifierServer(flags map[string]interface{}) {

	err := prepareVerifierServer(flags)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	setVerifierRoutes(r)

	// Run server
	err = r.Run(":" + PORT)
	if err != nil {
		panic(err)
	}
}

func prepareVerifierServer(flags map[string]interface{}) error {

	// Get keys from flags
	flagKeys := flags["keys"].(map[string]string)
	keys, err := getKeys(flagKeys, "public")
	if err != nil {
		return err
	}

	controllers.NewBaseController(&keys)
	return nil
}

func setVerifierRoutes(r *gin.Engine) {

	setBaseRoutes(r)
	// Verify endpoint
	endPoint := BASE_ROUTE + "/verify/:keyName"
	r.POST(endPoint, func(c *gin.Context) {
		controllers.VerifyController(c)
	})

}

func setBaseRoutes(r *gin.Engine) {

	// Health check endpoint
	endPoint := "/liveness"
	r.GET(endPoint, func(c *gin.Context) {
		controllers.HealthController(c)
	})

	// Swagger route
	endPoint = BASE_ROUTE + "/swagger/*any"
	r.GET(endPoint, ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func getKeys(flagKeys map[string]string, keyType string) (map[string]interface{}, error) {
	keys := map[string]interface{}{}

	// Set Engine
	crypto.Init()

	for keyFlagName, keyPemPath := range flagKeys {
		keyPem, err := crypto.GetKeyPem(&keyPemPath)
		if err != nil {
			return nil, err
		}

		var key interface{}

		if keyType == "private" {
			key, err = openssl.LoadPrivateKeyFromPEM(keyPem)
		} else {
			key, err = openssl.LoadPublicKeyFromPEM(keyPem)
		}

		if err != nil {
			return nil, err
		}

		// Set regular field, e.g. k1, k2
		keys[keyFlagName] = key

		// Set hash field
		setHashField(&keys, keyPem, key)

	}

	return keys, nil
}

func setHashField(keys *map[string]interface{}, keyPem []byte, key interface{}) {
	h := sha256.New()
	h.Write(keyPem)
	sha := hex.EncodeToString(h.Sum(nil))
	(*keys)[sha] = key
}

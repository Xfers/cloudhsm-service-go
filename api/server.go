package api

import (
	"github.com/Xfers/cloudhsm-service-go/api/controllers"
	"github.com/Xfers/cloudhsm-service-go/crypto"
	_ "github.com/Xfers/cloudhsm-service-go/docs"
	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-openssl"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title CloudHSM service API
// @version 1.0
// @description Standalone cloudhsm-service

// @host localhost:3000
// @BasePath /
// @schemes http
func RunSignerServer(flags map[string]interface{}) {

	err := prepareSignerServer(flags)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	setSignerRoutes(r)

	// Run server
	err = r.Run(":3000")
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

	//TODO: Set time limit in context
	// Digest endpoint
	r.POST("/digest", func(c *gin.Context) {
		controllers.DigestController(c)
	})
	// Sign endpoint
	r.POST("/sign/:keyName", func(c *gin.Context) {
		controllers.SignController(c)
	})

	// Pure Sign endpoint
	r.POST("/pure-sign/:keyName", func(c *gin.Context) {
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
	err = r.Run(":3000")
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
	r.POST("/verify/:keyName", func(c *gin.Context) {
		controllers.VerifyController(c)
	})

}

func setBaseRoutes(r *gin.Engine) {

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		controllers.HealthController(c)
	})

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func getKeys(flagKeys map[string]string, keyType string) (map[string]interface{}, error) {
	keys := map[string]interface{}{}
	for keyFlagName, keyPemPath := range flagKeys {
		keyPem, err := crypto.GetKeyPem(&keyPemPath)
		if err != nil {
			return nil, err
		}

		if keyType == "private" {
			keys[keyFlagName], err = openssl.LoadPrivateKeyFromPEM(keyPem)
		} else {
			keys[keyFlagName], err = openssl.LoadPublicKeyFromPEM(keyPem)
		}

		if err != nil {
			return nil, err
		}
	}

	return keys, nil
}

//TODO: Move these tests to test folder
// func test() {
// 	// Set test keys and data
// 	privatekey_pem, publickey_pem := generateTestKeys()
// 	data := "test data that is to be signed"

// 	// Digest
// 	d, err := crypto.Digest(data)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("Digest: ", d)
// 	fmt.Println("")
// 	//print line
// 	printLine()

// 	// Set Signers
// 	signer := crypto.NewSigner(privatekey_pem, d)                   //default, using go's package
// 	signerOpenSSL := crypto.NewSigner(privatekey_pem, d, "openssl") //openssl

// 	// Sign
// 	signature, err := signer.Sign()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Signature: ", signature)
// 	fmt.Println("")

// 	signatureOpenSSl, err := signerOpenSSL.Sign()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Signature OpenSSL: ", signatureOpenSSl)
// 	fmt.Println("")

// 	printLine()

// 	// Test Verify correct public key
// 	verifier := crypto.NewVerifier(publickey_pem, signature, data)
// 	verifierOpenSSL := crypto.NewVerifier(publickey_pem, signatureOpenSSl, data, "openssl")

// 	msg := "Verify correct public key: "

// 	result := verifier.Verify()
// 	printResult(msg, result)
// 	fmt.Println("")

// 	result = verifierOpenSSL.Verify()
// 	printResult("OpenSSL "+msg, result)
// 	fmt.Println("")

// 	printLine()

// 	// Test Verify wrong public key
// 	wrongPubKeyVerifier := crypto.NewVerifier("wrong public key", signature, data)
// 	wrongPubKeyVerifierOpenSSL := crypto.NewVerifier("wrong public key", signatureOpenSSl, data, "openssl")

// 	msg = "Result for wrong public key:"

// 	result = wrongPubKeyVerifier.Verify()
// 	printResult(msg, result)
// 	fmt.Println("")

// 	result = wrongPubKeyVerifierOpenSSL.Verify()
// 	printResult("OpenSSL "+msg, result)
// 	fmt.Println("")

// 	printLine()

// 	// Test Verify wrong data
// 	wrongDataVerifier := crypto.NewVerifier(publickey_pem, signature, "wrong data")
// 	wrongDataVerifierOpenSSL := crypto.NewVerifier(publickey_pem, signatureOpenSSl, "wrong data", "openssl")

// 	msg = "Result with wrong data:"

// 	result = wrongDataVerifier.Verify()
// 	printResult(msg, result)
// 	fmt.Println("")

// 	result = wrongDataVerifierOpenSSL.Verify()
// 	printResult("OpenSSL "+msg, result)
// 	fmt.Println("")

// 	printLine()

// 	// Test Verify wrong signature
// 	wrongSignatureVerifier := crypto.NewVerifier(publickey_pem, "wrong signature", data)
// 	wrongSignatureVerifierOpenSSL := crypto.NewVerifier(publickey_pem, "wrong signature", data, "openssl")

// 	msg = "Result with wrong signature:"

// 	result = wrongSignatureVerifier.Verify()
// 	printResult(msg, result)
// 	fmt.Println("")

// 	result = wrongSignatureVerifierOpenSSL.Verify()
// 	printResult("OpenSSL "+msg, result)
// 	fmt.Println("")

// 	printLine()
// }
// func generateTestKeys() (string, string) {
// 	// generate key
// 	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
// 	if err != nil {
// 		fmt.Printf("Cannot generate RSA key: %v", err)
// 		os.Exit(1)
// 	}
// 	publickey := &privatekey.PublicKey

// 	// dump private key to string
// 	privatekey_pem := string(pem.EncodeToMemory(&pem.Block{
// 		Type:  "RSA PRIVATE KEY",
// 		Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
// 	}))

// 	// dump public key to string
// 	publickey_pem := string(pem.EncodeToMemory(&pem.Block{
// 		Type:  "RSA PUBLIC KEY",
// 		Bytes: x509.MarshalPKCS1PublicKey(publickey),
// 	}))

// 	//print the keys
// 	fmt.Println("Private Key: ", privatekey_pem)
// 	fmt.Println("Public Key: ", publickey_pem)
// 	printLine()

// 	return privatekey_pem, publickey_pem
// }

// func printLine() {
// 	fmt.Println("------------------------------------------------------------")
// }

// func printResult(msg string, result bool) {
// 	if result {
// 		fmt.Println(msg, "Result: Success")
// 	} else {
// 		fmt.Println(msg, "Result: Fail")
// 	}
// }

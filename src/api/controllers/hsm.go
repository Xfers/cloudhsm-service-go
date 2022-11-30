package controllers

import (
	"net/http"

	"github.com/Xfers/cloudhsm-service-go/crypto"
	"github.com/Xfers/go-openssl"
	"github.com/gin-gonic/gin"
)

type DigestRequest struct {
	Data string `json:"data"`
}

type DigestResponse struct {
	Digest string `json:"digest"`
}

// Digest godoc
// @Summary Digest the data.
// @Description digest the data currently using sha256
// @Tags digest
// @Accept json
// @Produce json
// @Param data body DigestRequest true "Data to be digested"
// @Success 200 {object} DigestResponse
// @Router /api/digest [post]
func DigestController(c *gin.Context) {
	// Get data from request body json
	var digestRequest DigestRequest
	if err := c.ShouldBindJSON(&digestRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Digest
	d, err := crypto.Digest(digestRequest.Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return digest
	digestResponse := DigestResponse{
		Digest: d,
	}
	c.JSON(http.StatusOK, digestResponse)
}

type SignResponse struct {
	Result string `json:"result"`
}

// Sign godoc
// @Summary Sign the digest.
// @Description sign the digest using openssl or cloudhsm
// @Tags sign
// @Accept json
// @Produce json
// @Param keyName path string true "Key Name" example(k1) minLength(2) maxLength(2) pattern([a-z]+) style(simple) allowEmptyValue(false)
// @Success 200 {object} SignResponse
// @Router /api/sign/{keyName} [post]
func SignController(c *gin.Context) {

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get Key Name from path
	keyName := c.Param("keyName")
	if keyName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "keyName not provided"})
		return
	}

	// Get key from Base Controller Singleton
	baseController := GetBaseController()
	// check if set
	if baseController == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Base Controller not set"})
		return
	}
	priv := baseController.getKey(keyName).(openssl.PrivateKey)

	// Sign
	signer := crypto.NewSigner(&priv, string(body))
	signature, err := signer.Sign()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return signature
	signResponse := SignResponse{
		Result: signature,
	}
	c.JSON(http.StatusOK, signResponse)
}

type PureSignResponse struct {
	Result string `json:"result"`
}

// PureSign godoc
// @Summary Sign the data.
// @Description sign the data using openssl or cloudhsm
// @Tags pure-sign
// @Accept json
// @Produce json
// @Param keyName path string true "Key Name" example(k1) minLength(2) maxLength(2) pattern([a-z]+) style(simple) allowEmptyValue(false)
// @Success 200 {object} PureSignResponse
// @Router /api/pure-sign/{keyName} [post]
func PureSignController(c *gin.Context) {

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get Key Name from path
	keyName := c.Param("keyName")
	if keyName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "keyName not provided"})
		return
	}

	// Get key from Base Controller Singleton
	baseController := GetBaseController()
	// check if set
	if baseController == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Base Controller not set"})
		return
	}
	priv := baseController.getKey(keyName).(openssl.PrivateKey)

	digest := string(body)

	// Pure Sign
	signer := crypto.NewPureSigner(&priv, digest)
	signature, err := signer.Sign()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return signature
	pureSignResponse := PureSignResponse{
		Result: signature,
	}
	c.JSON(http.StatusOK, pureSignResponse)
}

type VerifyRequest struct {
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

type VerifyResponse struct {
	Valid bool `json:"valid"`
}

// Verify godoc
// @Summary Verify the Data.
// @Description verify the data using provided signature and public key, using openssl or cloudhsm
// @Tags verify
// @Accept json
// @Produce json
// @Param data body VerifyRequest true "Data to be verified"
// @Param keyName path string true "Key Name" example(k1) minLength(2) maxLength(2) pattern([a-z]+) style(simple) allowEmptyValue(false)
// @Success 200 {object} VerifyResponse
// @Router /api/verify/{keyName} [post]
func VerifyController(c *gin.Context) {
	// Get data from request body json
	var verifyRequest VerifyRequest
	if err := c.ShouldBindJSON(&verifyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get Key Name from path
	keyName := c.Param("keyName")
	if keyName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "keyName not provided"})
		return
	}

	// Get private key from Base Controller Singleton
	baseController := GetBaseController()
	// check if set
	if baseController == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Base Controller not set"})
		return
	}
	pub := baseController.getKey(keyName).(openssl.PublicKey)

	// Verify
	verifier := crypto.NewVerifier(&pub, verifyRequest.Signature, verifyRequest.Data)
	verified := verifier.Verify()

	// Return verification result
	verifyResponse := VerifyResponse{
		Valid: verified,
	}
	c.JSON(http.StatusOK, verifyResponse)
}

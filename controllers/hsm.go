package controllers

import (
	"net/http"

	"github.com/Xfers/cloudhsm-service-go/crypto"
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
// @Router /digest [post]
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

type SignRequest struct {
	Digest string `json:"digest"`
}

type SignResponse struct {
	Signature string `json:"signature"`
}

// Sign godoc
// @Summary Sign the digest.
// @Description sign the digest using openssl or cloudhsm
// @Tags sign
// @Accept json
// @Produce json
// @Param digest body SignRequest true "Digest to be signed"
// @Success 200 {object} SignResponse
// @Router /sign [post]
func SignController(c *gin.Context) {
	// Get data from request body json
	var signRequest SignRequest
	if err := c.ShouldBindJSON(&signRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Sign
	signer := crypto.NewSigner(c.GetString("private_key_path"), signRequest.Digest)
	signature, err := signer.Sign()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return signature
	signResponse := SignResponse{
		Signature: signature,
	}
	c.JSON(http.StatusOK, signResponse)
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
// @Success 200 {object} VerifyResponse
// @Router /verify [post]
func VerifyController(c *gin.Context) {
	// Get data from request body json
	var verifyRequest VerifyRequest
	if err := c.ShouldBindJSON(&verifyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify
	verifier := crypto.NewVerifier(c.GetString("public_key_path"), verifyRequest.Signature, verifyRequest.Data)
	verified := verifier.Verify()

	// Return verification result
	verifyResponse := VerifyResponse{
		Valid: verified,
	}
	c.JSON(http.StatusOK, verifyResponse)
}

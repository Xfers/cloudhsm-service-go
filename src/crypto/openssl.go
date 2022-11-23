package crypto

import (
	"encoding/base64"

	"github.com/Xfers/go-openssl"
)

type opensslSigner struct {
	priv *openssl.PrivateKey
	data string
}

func (s *opensslSigner) Sign() (string, error) {

	// // Get Digest
	digest, err := Digest(s.data)
	if err != nil {
		return "", err
	}

	sig, err := sign(digest, s.priv)
	if err != nil {
		return "", err
	}

	//return signature
	return base64.StdEncoding.EncodeToString(sig), nil
}

type opensslPureSigner struct {
	priv   *openssl.PrivateKey
	digest string
}

func (s *opensslPureSigner) Sign() (string, error) {

	sig, err := sign(s.digest, s.priv)
	if err != nil {
		return "", err
	}

	//return signature base64 encoded
	return base64.StdEncoding.EncodeToString(sig), nil
}

func sign(digest string, priv *openssl.PrivateKey) ([]byte, error) {
	digestBa, err := base64.StdEncoding.DecodeString(digest)
	if err != nil {
		return nil, err
	}

	// Sign
	sig, err := (*priv).Sign(openssl.SHA256_Method, digestBa)
	if err != nil {
		return nil, err
	}

	//return signature
	return sig, nil
}

type opensslVerifier struct {
	pub       *openssl.PublicKey
	signature string
	data      string
}

func (v *opensslVerifier) Verify() bool {

	digest, err := Digest(v.data)
	if err != nil {
		return false
	}
	digestBa, err := base64.StdEncoding.DecodeString(digest)
	if err != nil {
		return false
	}

	sig, err := base64.StdEncoding.DecodeString(v.signature)
	if err != nil {
		return false
	}

	//verify
	err = (*v.pub).VerifyPKCS1v15(openssl.SHA256_Method, digestBa, sig)

	return err == nil
}

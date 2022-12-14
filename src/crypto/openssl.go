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

	// Sign
	sig, err := (*s.priv).Sign(openssl.SHA256_Method, []byte(s.data))
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

	// determine if the digest is base64 encoded and should be decoded (coming from web api in our case)
	digestBa, err := base64.StdEncoding.DecodeString(s.digest)
	if err != nil {
		digestBa = []byte(s.digest)
	}

	sig, err := (*s.priv).PureSign(openssl.SHA256_Method, digestBa)

	if err != nil {
		return "", err
	}

	//return signature base64 encoded
	return base64.StdEncoding.EncodeToString(sig), nil
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

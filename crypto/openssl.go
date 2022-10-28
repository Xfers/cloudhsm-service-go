package crypto

import (
	"encoding/base64"

	openssl "github.com/libp2p/go-openssl"
)

type opensslSigner struct {
	privateKeyPem string
	digest        string
}

func (s *opensslSigner) Sign() (string, error) {

	priv, err := openssl.LoadPrivateKeyFromPEM([]byte(s.privateKeyPem))
	if err != nil {
		return "", err
	}

	// get []byte from digest
	digestBa, err := base64.URLEncoding.DecodeString(s.digest)
	if err != nil {
		return "", err
	}

	//sign
	sig, err := priv.SignPKCS1v15(openssl.SHA256_Method, digestBa)
	if err != nil {
		return "", err
	}

	//return signature
	return base64.URLEncoding.EncodeToString(sig), nil
}

type opensslVerifier struct {
	publicKeyPem string
	signature    string
	data         string
}

func (v *opensslVerifier) Verify() bool {

	//get public key from string
	priv, err := openssl.LoadPublicKeyFromPEM([]byte(v.publicKeyPem))
	if err != nil {
		return false
	}

	//get []byte from data
	digest, err := Digest(v.data)
	if err != nil {
		return false
	}
	digestBa, err := base64.URLEncoding.DecodeString(digest)
	if err != nil {
		return false
	}

	//get []byte from signature
	sig, err := base64.URLEncoding.DecodeString(v.signature)
	if err != nil {
		return false
	}

	//verify
	err = priv.VerifyPKCS1v15(openssl.SHA256_Method, digestBa, sig)

	return err == nil
}

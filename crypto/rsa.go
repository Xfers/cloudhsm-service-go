package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type rsaSigner struct {
	privateKeyPem []byte
	digest        string
}

func (s *rsaSigner) Sign() (string, error) {

	// get private key from string
	priv, err := getPrivateKeyFromPem(s.privateKeyPem)
	if err != nil {
		return "", err
	}

	// get []byte from digest
	digestBa, err := base64.URLEncoding.DecodeString(s.digest)
	if err != nil {
		return "", err
	}

	// sign
	sig, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, digestBa, nil)
	if err != nil {
		return "", err
	}

	//return signature
	return base64.URLEncoding.EncodeToString(sig), nil
}

type verifier struct {
	publicKeyPem []byte
	signature    string
	data         string
}

func (v *verifier) Verify() bool {

	//get public key from string
	pub, err := getPublicKeyFromPem(v.publicKeyPem)
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
	err = rsa.VerifyPSS(pub, crypto.SHA256, digestBa, sig, nil)

	return err == nil
}

func getPrivateKeyFromPem(pemPrivateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		return nil, errors.New("failed to decode Private Key")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func getPublicKeyFromPem(pemPublicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemPublicKey)
	if block == nil {
		return nil, errors.New("failed to decode Public Key")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return pub, nil
}

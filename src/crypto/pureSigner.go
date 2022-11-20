package crypto

import (
	openssl "github.com/libp2p/go-openssl"
)

type PureSigner interface {
	Sign() (string, error)
}

func NewPureSigner(priv *openssl.PrivateKey, digest string) Signer {

	Init()
	return &opensslPureSigner{priv, digest}
}

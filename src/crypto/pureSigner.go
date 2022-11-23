package crypto

import (
	"github.com/Xfers/go-openssl"
)

type PureSigner interface {
	Sign() (string, error)
}

func NewPureSigner(priv *openssl.PrivateKey, digest string) Signer {

	Init()
	return &opensslPureSigner{priv, digest}
}

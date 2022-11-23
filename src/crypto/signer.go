package crypto

import (
	"github.com/Xfers/go-openssl"
)

type Signer interface {
	Sign() (string, error)
}

func NewSigner(priv *openssl.PrivateKey, data string) Signer {

	return &opensslSigner{priv, data}

}

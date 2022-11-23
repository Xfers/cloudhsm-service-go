package crypto

import (
	"github.com/Xfers/go-openssl"
)

type Verifier interface {
	Verify() bool
}

func NewVerifier(publicKey *openssl.PublicKey, signature, data string) Verifier {

	Init()
	return &opensslVerifier{publicKey, signature, data}

}

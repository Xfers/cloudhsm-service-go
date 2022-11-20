package crypto

import (
	openssl "github.com/libp2p/go-openssl"
)

type Verifier interface {
	Verify() bool
}

func NewVerifier(publicKey *openssl.PublicKey, signature, data string) Verifier {

	Init()
	return &opensslVerifier{publicKey, signature, data}

}

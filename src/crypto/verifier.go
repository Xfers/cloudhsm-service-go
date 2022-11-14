package crypto

import openssl "github.com/libp2p/go-openssl"

type Verifier interface {
	Verify() bool
}

func NewVerifier(publicKey *openssl.PublicKey, signature, data string) Verifier {

	//default CloudHSM
	m := "CloudHSM"

	// If CloudHSM not reachable, use OpenSSL
	if !IsCloudHSMReachable() {
		m = "openssl"
	}

	switch m {
	case "openssl":
		return &opensslVerifier{publicKey, signature, data}
	default:
		//TODO: implement CloudHSM
		return nil
	}
}

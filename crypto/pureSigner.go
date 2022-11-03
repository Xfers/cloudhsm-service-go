package crypto

import openssl "github.com/libp2p/go-openssl"

type PureSigner interface {
	Sign() (string, error)
}

func NewPureSigner(priv *openssl.PrivateKey, digest string) Signer {

	//default CloudHSM
	m := "CloudHSM"

	// If CloudHSM not reachable, use OpenSSL
	if !IsCloudHSMReachable() {
		m = "openssl"
	}

	switch m {
	case "openssl":
		return &opensslPureSigner{priv, digest}
	default:
		//TODO: implement CloudHSM
		return nil
	}
}

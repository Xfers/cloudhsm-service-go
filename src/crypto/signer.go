package crypto

import openssl "github.com/libp2p/go-openssl"

type Signer interface {
	Sign() (string, error)
}

func NewSigner(priv *openssl.PrivateKey, data string) Signer {

	//default CloudHSM
	m := "CloudHSM"

	// If CloudHSM not reachable, use OpenSSL
	if !IsCloudHSMReachable() {
		m = "openssl"
	}

	switch m {
	case "openssl":
		return &opensslSigner{priv, data}
	default:
		//TODO: implement CloudHSM
		return nil
	}
}

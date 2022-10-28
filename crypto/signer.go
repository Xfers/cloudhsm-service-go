package crypto

type Signer interface {
	Sign() (string, error)
}

func NewSigner(privateKeyPem, digest string, mode ...string) Signer {

	m := "rsa"
	if len(mode) > 0 {
		m = mode[0]
	}

	switch m {
	case "openssl":
		return &opensslSigner{privateKeyPem, digest}
	default:
		return &signer{privateKeyPem, digest}
	}
}

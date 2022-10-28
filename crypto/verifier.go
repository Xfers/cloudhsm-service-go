package crypto

type Verifier interface {
	Verify() bool
}

func NewVerifier(publicKeyPem, signature, data string, mode ...string) Verifier {

	m := "rsa"
	if len(mode) > 0 {
		m = mode[0]
	}

	switch m {
	case "openssl":
		return &opensslVerifier{publicKeyPem, signature, data}
	default:
		return &verifier{publicKeyPem, signature, data}
	}
}

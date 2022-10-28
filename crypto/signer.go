package crypto

type Signer interface {
	Sign() (string, error)
}

func NewSigner(privateKeyPemPath, digest string) Signer {

	// Get key from pem file
	privateKeyPem, err := GetKeyPem(privateKeyPemPath)
	if err != nil {
		return nil
	}

	//default CloudHSM
	m := "CloudHSM"

	// If CloudHSM not reachable, use OpenSSL
	if !IsCloudHSMReachable() {
		m = "openssl"
	}

	switch m {
	case "openssl":
		return &opensslSigner{privateKeyPem, digest}
	default:
		// Here pretend to be CloudHSM but it is golang's rsa
		return &rsaSigner{privateKeyPem, digest}
	}
}

package crypto

type Verifier interface {
	Verify() bool
}

func NewVerifier(publicKeyPemPath, signature, data string) Verifier {

	// Get public key from pem file
	publicKeyPem, err := GetKeyPem(publicKeyPemPath)
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
		return &opensslVerifier{publicKeyPem, signature, data}
	default:
		// Here pretend to be CloudHSM but it is golang's rsa
		return &verifier{publicKeyPem, signature, data}
	}
}

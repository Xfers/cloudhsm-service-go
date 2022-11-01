package crypto

import "io/ioutil"

// Temporary solution to use OpenSSL instead of CloudHSM
func IsCloudHSMReachable() bool {
	return false
}

func GetKeyPem(KeyPath *string) ([]byte, error) {
	// Read key from pem file
	keyPem, err := ioutil.ReadFile(*KeyPath)
	if err != nil {
		return nil, err
	}
	return keyPem, nil
}

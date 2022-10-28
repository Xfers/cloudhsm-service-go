package crypto

import "io/ioutil"

// Temporary solution to use OpenSSL instead of CloudHSM
func IsCloudHSMReachable() bool {
	return false
}

func GetKeyPem(privateKeyPemPath string) ([]byte, error) {
	// Read key from pem file
	privateKeyPem, err := ioutil.ReadFile(privateKeyPemPath)
	if err != nil {
		return nil, err
	}
	return privateKeyPem, nil
}

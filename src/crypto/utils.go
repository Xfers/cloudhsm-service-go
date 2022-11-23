package crypto

import (
	"io/ioutil"
)

func GetKeyPem(KeyPath *string) ([]byte, error) {
	// Read key from pem file
	keyPem, err := ioutil.ReadFile(*KeyPath)
	if err != nil {
		return nil, err
	}
	return keyPem, nil
}

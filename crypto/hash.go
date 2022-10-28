package crypto

import (
	"crypto/sha256"
	"encoding/base64"
)

func Digest(data string) (string, error) {
	ba := []byte(data)
	hasher := sha256.New()
	_, err := hasher.Write(ba)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)), nil
}

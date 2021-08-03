package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomString generates a random string with p bytes of length.
// Specifying 10 in the p parameter will result in the length of 20.
func RandomString(p int) (string, error) {
	if p <= 0 {
		p = 10
	}
	arr := make([]byte, p)
	_, err := rand.Read(arr)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(arr), nil
}

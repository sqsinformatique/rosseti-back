package crypto

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
)

func HashString(data string) string {
	if data == "" {
		return ""
	}

	h := sha256.New()

	// Write in Hash interface never returns an error.
	// nolint
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

var (
	ErrMismatchedHashAndPassword = errors.New("hashedPassword is not the hash of the given password")
)

func CompareHash(hashedPassword, password string) error {
	otherP := HashString(password)

	if subtle.ConstantTimeCompare([]byte(hashedPassword), []byte(otherP)) == 1 {
		return nil
	}

	return ErrMismatchedHashAndPassword
}

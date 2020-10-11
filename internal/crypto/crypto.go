package crypto

import (
	"crypto/sha256"
	"encoding/hex"
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

package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/subtle"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
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

func GenerateSign() (*rsa.PrivateKey, error) {
	// The GenerateKey method takes in a reader that returns random bits, and
	// the number of bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func MarshalSign(key *rsa.PrivateKey) (privateKey, publicKey string) {
	privateKeyData := x509.MarshalPKCS1PrivateKey(key)
	publicKeyData := x509.MarshalPKCS1PublicKey(&key.PublicKey)

	return string(privateKeyData), string(publicKeyData)
}

func UnmarshalPrivate(privateKey string) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey([]byte(privateKey))
}

func UnmarshalPublic(publicKey string) (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey([]byte(publicKey))
}

func DataSign(data interface{}, key *rsa.PrivateKey) (string, error) {
	d, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	dataHash := sha256.New()
	_, err = dataHash.Write(d)
	if err != nil {
		return "", err
	}
	dataHashSum := dataHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, dataHashSum, nil)
	if err != nil {
		return "", err
	}

	return string(signature), nil
}

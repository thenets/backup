package secret

import (
	"crypto/sha512"
	"encoding/base64"
)

// Spec holds the ssh priv and pub keys
type Spec struct {
	PrivateKey string
	PublicKey  string
}

// GetPrivateKeyHash returns sha512 hash format
func (s *Spec) GetPrivateKeyHash() string {
	return bytesToHash([]byte(s.PrivateKey))
}

// GetPublicKeyHash returns sha512 hash format
func (s *Spec) GetPublicKeyHash() string {
	return bytesToHash([]byte(s.PublicKey))
}

// bytesToHash generates sha512 hash from 'b', encode to base64,
// and returns as string format.
// Also returns nothing is 'b' is empty.
func bytesToHash(b []byte) string {
	// Check if empty
	if len(b) == 0 {
		return ""
	}

	// Generate hash in byte format
	var sha512Instance = sha512.New()
	sha512Instance.Write([]byte(b))
	var hash = sha512Instance.Sum(nil)

	// Encode to base64 string format
	var hashBase64 = base64.StdEncoding.EncodeToString(hash)

	return hashBase64
}

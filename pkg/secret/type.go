package secret

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

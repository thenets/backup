package secret

// Secret holds secret kind info
type Secret struct {
	Kind     string
	Metadata struct {
		ID   string
		Name string
	}

	// Spec holds the ssh priv and pub keys
	Spec struct {
		PrivateKey string
		PublicKey  string
	}
}

// GetPrivateKeyHash returns sha512 hash format
func (s *Secret) GetPrivateKeyHash() string {
	return bytesToHash([]byte(s.Spec.PrivateKey))
}

// GetPublicKeyHash returns sha512 hash format
func (s *Secret) GetPublicKeyHash() string {
	return bytesToHash([]byte(s.Spec.PublicKey))
}

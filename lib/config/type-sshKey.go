package config

// SSHKeyData holds the ssh priv and pub keys
type SSHKeyData struct {
	Kind string
	Metadata
	spec struct {
		PrivateKey string
		PublicKey  string
	}
}

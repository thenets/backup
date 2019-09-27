package main

// SSHKeyData holds the ssh priv and pub keys
type SSHKeyData struct {
	Kind string
	Metadata struct {
		ID string
		Name string
	}
	spec struct {
		PrivateKey string
		PublicKey string
	}
}
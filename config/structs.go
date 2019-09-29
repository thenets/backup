package config

// File holds a config file content of any type
type File struct {
	Kind     string
	filePath string

	rsync  RsyncData
	sshKey SSHKeyData
}

// Metadata is the base of any config file
type Metadata struct {
	ID   string
	Name string
}

// RsyncData data struct for rsync config file
type RsyncData struct {
	Kind string
	Metadata
	Spec struct {
		RemoteDir           string
		LocalDestinationDir string
		CustomArgs          string
		Server              struct {
			Host         string
			Port         int
			Password     string
			PrivateKeyID string
		}
	}
}

// SSHKeyData holds the ssh priv and pub keys
type SSHKeyData struct {
	Kind string
	Metadata
	spec struct {
		PrivateKey string
		PublicKey  string
	}
}

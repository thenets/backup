package config

// File holds a config file content of any type
type File struct {
	Kind     string
	filePath string

	ssh    SSHData
	sshKey SSHKeyData
}

// Metadata is the base of any config file
type Metadata struct {
	ID   string
	Name string
}

// SSHData data struct for ssh config file
type SSHData struct {
	Kind string
	Metadata
	Spec struct {
		RemoteDir           string
		LocalDestinationDir string
		CustomRsyncArgs     string
		CustomSSHArgs       string
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

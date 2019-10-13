package ssh

// Spec data struct for ssh config file
type Spec struct {
	RemoteDir           string
	LocalDestinationDir string
	CustomRsyncArgs     string
	CustomSSHArgs       string

	Server struct {
		Host         string
		Port         int
		Password     string
		PrivateKeyID string
	}
}

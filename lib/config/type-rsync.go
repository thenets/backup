package config

// RsyncData data struct for rsync config file
type RsyncData struct {
	Kind string
	Metadata
	Spec struct {
		RemoteDir string
		LocalDestinationDir string
		CustomArgs string
		Server struct {
			Host string
			Port int
			Password string
			PrivateKeyID string
		}
	}
}

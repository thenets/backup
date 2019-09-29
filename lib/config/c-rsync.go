package config

// RsyncData data struct for rsync config file
type RsyncData struct {
	Kind string
	Metadata struct {
		ID string
		name string
	}
	Spec struct {
		Server struct {
			Host string
			Port int
			Password string
			PrivateKeyID string
		}
		RemoteDir string
		LocalDestinationDir string
		CustomArgs string
	}
}

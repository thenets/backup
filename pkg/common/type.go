package common

import "github.com/thenets/backup/pkg/ssh"

// All  the struct is based on Kubernetes.
// More info at:
// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds

// File holds a config file content of any type
type File struct {
	Kind     string
	filePath string

	ssh    ssh.Data
	sshKey ssh.KeyData
}

// Metadata is the base of any config file
type Metadata struct {
	ID   string
	Name string
}

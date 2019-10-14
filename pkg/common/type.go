package common

import (
	"github.com/spf13/viper"
	"github.com/thenets/backup/utils"
)

// All  the struct is based on Kubernetes.
// More info at:
// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds

// Loads makes coffee TODO
func Loads(filePath string) {
	viper.SetConfigName("./secrets/key")

	err := viper.ReadInConfig()
	utils.Check(err, "can't read the config file")
}

// File holds a config file content of any type
type File struct {
	// Original raw config file content
	raw []byte
}

// GetMetadata returns file's metadata type
func (File) GetMetadata() Metadata {
	var metadata Metadata

	return metadata
}

// Metadata is the base of any config file
type Metadata struct {
	ID   string
	Name string
}

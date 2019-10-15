package common

import (
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"github.com/thenets/backup/pkg/secret"
	"github.com/thenets/backup/pkg/ssh"
	"github.com/thenets/backup/utils"
)

// All  the struct is based on Kubernetes.
// More info at:
// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds

// Config holds a config file content of any type
type Config struct {
	// Original viper instance created from the raw config file
	viperInstance *viper.Viper
}

// Loads makes coffee TODO
func Loads(filePath string) Config {
	// TODO validates file exist

	var viperInstance = viper.New()

	// Read 'filePath'
	path, fileName := utils.SplitPathAndFileName(filePath)
	fileNameWithoutExtension := utils.RemoveFileNameExtension(fileName)

	viperInstance.SetConfigName(fileNameWithoutExtension)
	viperInstance.AddConfigPath(path)

	err := viperInstance.ReadInConfig()
	utils.Check(err, "can't read the config file")

	// TODO validate keys

	// Populate config
	var c Config
	c.viperInstance = viperInstance

	return c
}

// GetViper returns viper instance
func (c *Config) GetViper() *viper.Viper {
	return c.viperInstance
}

// KindSecret returns 'secret.Secret' if config is Secret kind
// or panic if isn't.
func (c *Config) KindSecret() secret.Secret {
	// Validation
	if strings.ToLower(c.GetViper().GetString("kind")) != "sshkey" {
		var e = errors.New("is not a 'secret.Spec' type")
		utils.Check(e, "config file invalid")
	}

	var s secret.Secret
	s.Spec.PrivateKey = c.GetViper().GetString("spec.privatekey")
	s.Spec.PublicKey = c.GetViper().GetString("spec.publickey")

	return s
}

// SSH returns 'ssh.SSH' if config is Secret kind
// or panic if isn't.
func (c *Config) SSH() ssh.SSH {
	var err error

	// Validation
	if strings.ToLower(c.GetViper().GetString("kind")) != "ssh" {
		var e = errors.New("is not a 'ssh.Spec' type")
		utils.Check(e, "config file invalid")
	}

	// TODO improve to some unmarshal function to read config
	var s ssh.SSH
	s.Spec.RemoteDir = c.GetViper().GetString("spec.RemoteDir")
	s.Spec.LocalDestinationDir = c.GetViper().GetString("spec.LocalDestinationDir")
	s.Spec.CustomRsyncArgs = c.GetViper().GetString("spec.CustomRsyncArgs")
	s.Spec.CustomSSHArgs = c.GetViper().GetString("spec.CustomSSHArgs")
	s.Spec.Server.Host = c.GetViper().GetString("spec.server.Host")
	s.Spec.Server.Port, err = strconv.Atoi(c.GetViper().GetString("spec.server.Port"))
	utils.Check(err, "Can't convert 's.Spec.Server.Port' to int format")
	s.Spec.Server.Password = c.GetViper().GetString("spec.server.Password")
	s.Spec.Server.PrivateKeyID = c.GetViper().GetString("spec.server.PrivateKeyID")

	return s
}

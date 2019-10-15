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

	metadata struct {
		ID   string
		Name string
	}
}

// Loads makes coffee TODO
func Loads(filePath string) Config {
	var viperInstace = viper.New()

	// Read 'filePath'
	path, fileName := utils.SplitPathAndFileName(filePath)
	fileNameWithoutExtension := utils.RemoveFileNameExtension(fileName)

	viperInstace.SetConfigName(fileNameWithoutExtension)
	viperInstace.AddConfigPath(path)

	err := viperInstace.ReadInConfig()
	utils.Check(err, "can't read the config file")

	// TODO validate keys

	// Populate config
	var c Config
	c.viperInstance = viperInstace
	c.metadata.ID = viperInstace.GetString("id")
	c.metadata.Name = viperInstace.GetString("name")

	return c
}

// GetViper returns viper instance
func (c *Config) GetViper() *viper.Viper {
	return c.viperInstance
}

// SpecSecret returns 'secret.Spec' if config is Secret kind
// or panic if isn't.
func (c *Config) SpecSecret() secret.Spec {
	// Validation
	if strings.ToLower(c.GetViper().GetString("kind")) != "sshkey" {
		var e = errors.New("is not a 'secret.Spec' type")
		utils.Check(e, "config file invalid")
	}

	var s secret.Spec
	s.PrivateKey = c.GetViper().GetString("spec.privatekey")
	s.PublicKey = c.GetViper().GetString("spec.publickey")

	return s
}

// SpecSSH returns 'ssh.Spec' if config is Secret kind
// or panic if isn't.
func (c *Config) SpecSSH() ssh.Spec {
	var err error

	// Validation
	if strings.ToLower(c.GetViper().GetString("kind")) != "ssh" {
		var e = errors.New("is not a 'ssh.Spec' type")
		utils.Check(e, "config file invalid")
	}

	// TODO improve to some unmarshal function to read config
	var s ssh.Spec
	s.RemoteDir = c.GetViper().GetString("spec.RemoteDir")
	s.LocalDestinationDir = c.GetViper().GetString("spec.LocalDestinationDir")
	s.CustomRsyncArgs = c.GetViper().GetString("spec.CustomRsyncArgs")
	s.CustomSSHArgs = c.GetViper().GetString("spec.CustomSSHArgs")
	s.Server.Host = c.GetViper().GetString("spec.server.Host")
	s.Server.Port, err = strconv.Atoi(c.GetViper().GetString("spec.server.Port"))
	utils.Check(err, "Can't convert 's.Server.Port' to int format")
	s.Server.Password = c.GetViper().GetString("spec.server.Password")
	s.Server.PrivateKeyID = c.GetViper().GetString("spec.server.PrivateKeyID")

	return s
}

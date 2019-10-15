package common

import (
	"errors"
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

// Load makes coffee TODO
func Load(filePath string) Config {
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

func LoadSecret(filePath string) secret.Secret {
	secretConfig := Load(filePath)
	return secretConfig.KindSecret()
}

func LoadSSH(filePath string) ssh.SSH {
	sshConfig := Load(filePath)
	return sshConfig.KindSSH()
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
	c.GetViper().Unmarshal(&s)

	return s
}

// KindSSH returns 'ssh.SSH' if config is Secret kind
// or panic if isn't.
func (c *Config) KindSSH() ssh.SSH {
	var err error

	// Validation
	if strings.ToLower(c.GetViper().GetString("kind")) != "ssh" {
		var e = errors.New("is not a 'ssh.Spec' type")
		utils.Check(e, "config file invalid")
	}

	// Unmarshal
	var s ssh.SSH
	err = c.GetViper().Unmarshal(&s)
	utils.Check(err, "unable to unmarshal SSH config file")

	return s
}

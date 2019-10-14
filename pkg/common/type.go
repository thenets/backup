package common

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
	"github.com/thenets/backup/pkg/secret"
	"github.com/thenets/backup/utils"
)

// All  the struct is based on Kubernetes.
// More info at:
// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds

// Loads makes coffee TODO
func Loads(filePath string) Config {
	var viperInstace = viper.New()

	viperInstace.SetConfigName("key")
	viperInstace.AddConfigPath("./secrets/")

	err := viperInstace.ReadInConfig()
	utils.Check(err, "can't read the config file")

	var c Config
	c.viperInstance = viperInstace

	return c
}

// Config holds a config file content of any type
type Config struct {
	// Original viper instance created from the raw config file
	viperInstance *viper.Viper
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

// Metadata is the base of any config file
type Metadata struct {
	ID   string
	Name string
}

package config

import (
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// File holds a config file content of any type
type File struct {
	Kind     string
	filePath string

	rsync  RsyncData
	sshKey SSHKeyData
}

// Metadata is the base of any config file
type Metadata struct {
	ID   string
	Name string
}

// reload reloads the current config file
func (c File) reload() error {
	var err error

	return err
}

// load loads a config file
func (c *File) load(filePath string) error {
	var err error

	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Get config kind name
	err = yaml.Unmarshal([]byte(f), &c)
	if err != nil {
		return err
	}

	// Load data
	switch strings.ToLower(c.Kind) {
	case "rsync":
		var rsync RsyncData
		err = yaml.Unmarshal(f, &rsync)
		if err != nil {
			return err
		}
		c.rsync = rsync
	case "sshkey":
		var sshKey SSHKeyData
		err = yaml.Unmarshal(f, &sshKey)
		if err != nil {
			return err
		}
		c.sshKey = sshKey
	default:
		return errors.New("config kind not recognized")
	}

	c.filePath = filePath

	return err
}

// Rsync returns the RSync file content or error if file type is incorrect
func (c File) Rsync() (RsyncData, error) {
	var err error
	var rsync RsyncData

	if c.Kind != "rsync" {
		return rsync, errors.New("config is not a 'rsync' kind")
	}

	return c.rsync, err
}

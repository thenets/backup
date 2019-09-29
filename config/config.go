package config

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// Loads a config file
func Loads(filePath string) (File, error) {
	var configFile File

	err := configFile.load(filePath)

	return configFile, err

}

// reload reloads the current config file
func (c File) reload() error {
	var err error

	return err
}

// load loads a config file
func (c *File) load(filePath string) error {
	var err error

	f, err := readYamlFile(filePath)
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

// readYamlFile read a yaml file, fix some issues and return in []byte format
func readYamlFile(filePath string) ([]byte, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(nil), err
	}

	// Make all keys become lower case
	//var content = string(f)
	bufferReader := bufio.NewReader(bytes.NewReader(f))

	for {
		lineContent, isPrefix, err := bufferReader.ReadLine()

		// TODO replace and lowercase all keys here

		if isPrefix {
			return []byte(nil), errors.New("file is too big")
		}

		if err != nil {
			err = error(nil)
			break
		}
	}

	return f, err
}

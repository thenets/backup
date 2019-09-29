package config

import (
	"io/ioutil"
	"log"
)

// CheckConfigFileKind returns the config file kind
func CheckConfigFileKind(filePath string) (string, error) {
	var err error
	var kind string

	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return kind, err
	}
	content := string(f)

	log.Println(content)

	return kind, err
}

package helpers

import (
	"log"
	"os"
)

// IsDirectory returns true if is a directory
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// Check the error obj and panic if not nil
func Check(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Fatal(err)
	}
}

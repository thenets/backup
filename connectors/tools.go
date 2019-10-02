package connectors

import (
	"log"
	"os"
)

// isDirectory returns true if is a directory
func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func getLogsPath() string {
	var logsDir = "/tmp/tnb/logs/"

	if !isDirectory(logsDir) {
		os.MkdirAll(logsDir, 0755)
	}

	return logsDir
}

func getCacheDir() string {
	var cacheDir = "/tmp/tnb/cache/"

	if !isDirectory(cacheDir) {
		os.MkdirAll(cacheDir, 0755)
	}

	return cacheDir
}

func check(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Fatal(err)
	}
}

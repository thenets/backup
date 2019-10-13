package helpers

import "os"

//GetLogsPath returns the absolut logs path
func GetLogsPath() string {
	var logsDir = "/tmp/tnb/logs/"

	if !IsDirectory(logsDir) {
		os.MkdirAll(logsDir, 0755)
	}

	return logsDir
}

//GetCacheDir returns the absolut caches path
func GetCacheDir() string {
	var cacheDir = "/tmp/tnb/cache/"

	if !IsDirectory(cacheDir) {
		os.MkdirAll(cacheDir, 0755)
	}

	return cacheDir
}

package utils

import (
	"errors"
	"log"
	"os"
	"strings"
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
		//log.Fatal(err)
		panic(err)
	}
}

// StringReverse return 's' reverted
func StringReverse(s string) string {
	var out string
	for i := len(s) - 1; i >= 0; i-- {
		out += string(s[i])
	}
	return out
}

// SplitPathAndFileName returns 'path' and 'fileName' separated
func SplitPathAndFileName(filePath string) (string, string) {
	arr := strings.Split(filePath, "/")
	path := strings.Join(arr[:len(arr)-1], "/")
	fileName := arr[len(arr)-1]

	return path, fileName
}

// RemoveFileNameExtension removes latest content after "." from fileName
// and returns new string
func RemoveFileNameExtension(fileName string) string {
	sSlice := strings.Split(fileName, ".")
	withoutExtension := strings.Join(sSlice[:len(sSlice)-1], ".")

	return withoutExtension
}

// CreateCacheDir returns a temporary cache dir path
func CreateCacheDir(id string) (string, error) {
	var err error

	// Create cache dir
	var cacheDir = getCacheDir() + id + "/"
	if !IsDirectory(cacheDir) {
		os.MkdirAll(cacheDir, 0755)
	}
	if !IsDirectory(cacheDir) {
		err = errors.New("cache dir can't be created: " + cacheDir)
	}

	return cacheDir, err

}

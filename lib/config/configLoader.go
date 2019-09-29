package config

// Loads a config file
func Loads(filePath string) (File, error) {
	var configFile File

	err := configFile.load(filePath)

	return configFile, err

}

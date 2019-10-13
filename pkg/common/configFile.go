package common

import "github.com/spf13/viper"

// Loads makes coffee TODO
func Loads(filePath string) {
	viper.SetConfigName("./secrets/key")
}

package env

import (
	"fmt"

	"github.com/spf13/viper"
)

const keyAddress = "ADDRESS"
const keyTempDir = "TEMP_DIR"
const keyLibreOfficePath = "LIBREOFFICE_PATH"

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file was not found"))
		} else {
			panic(err)
		}
	}
	viper.AutomaticEnv()
}

func Validate() error {
	if !viper.IsSet(keyAddress) {
		return createKeyError(keyAddress)
	}
	if !viper.IsSet(keyTempDir) {
		return createKeyError(keyTempDir)
	}
	if !viper.IsSet(keyLibreOfficePath) {
		return createKeyError(keyLibreOfficePath)
	}
	return nil
}

func createKeyError(key string) error {
	return fmt.Errorf("env for key %s is missing", key)
}

func GetAddress() string {
	return viper.GetString(keyAddress)
}

func GetTempDir() string {
	return viper.GetString(keyTempDir)
}

func GetLibreOfficePath() string {
	return viper.GetString(keyLibreOfficePath)
}

package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// MustLoadConfig will panic if it cannot create client or parse the config/secret.
func MustLoadConfig(fPath, fname, ftype string, dst any) {
	if err := LoadConfig(fPath, fname, ftype, dst); err != nil {
		log.Panicf("cannot load configuration from %s: %s", fPath, err.Error())
	}
}

// Load loads configuration into dest
func LoadConfig(filePath, fileName, extension string, dest any) (err error) {
	switch extension {
	case "env":
		viper.SetConfigType("env")
	case "yaml":
		viper.SetConfigType("yaml")
	case "json":
		viper.SetConfigType("json")
	default:
		return fmt.Errorf("not support this file type: %s", extension)
	}
	viper.AddConfigPath(filePath)
	viper.SetConfigName(fileName)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&dest)
	return err
}

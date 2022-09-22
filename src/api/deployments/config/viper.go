package config

import (
	"github.com/spf13/viper"
)

func InitViper() error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cp2")
	return nil
}

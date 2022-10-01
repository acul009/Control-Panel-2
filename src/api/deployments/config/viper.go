package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cp2")
	fmt.Println("initializing viper")
	return nil
}

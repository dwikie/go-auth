package config

import (
	"github.com/spf13/viper"
)

var Cfg *config

func Init() {
	viper.AutomaticEnv()
	viper.SetConfigFile("config.yml")
	viper.ReadInConfig()

	viper.Unmarshal(&Cfg)
}

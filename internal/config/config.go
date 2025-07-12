package config

import "github.com/spf13/viper"

func LoadConfig(path string) (Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	return Config{}, nil
}

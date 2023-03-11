package config

import "github.com/spf13/viper"

type Config struct {
	Port int
}

func ReadConfigFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if error := viper.ReadInConfig(); error != nil {
		return Config{}, error
	}

	return Config{
			Port: viper.GetInt("port"),
		},
		nil
}

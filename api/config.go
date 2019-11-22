package api

import "github.com/spf13/viper"

type Config struct {
	// The port to bind the web application server to
	Port int
}

func InitConfig() (*Config, error) {
	config := &Config{
		Port:       viper.GetInt("Port"),
	}
	if config.Port == 0 {
		config.Port = 9092
	}
	return config, nil
}
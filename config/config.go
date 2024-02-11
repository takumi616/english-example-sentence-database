package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port	int	`env:"APP_CONTAINER_PORT"`
}

func GetConfig() (*Config, error) {
	config := &Config{}
	//Automatically set environment variables to struct field
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

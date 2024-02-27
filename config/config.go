package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port             int    `env:"APP_CONTAINER_PORT"`
	DatabaseHost     string `env:"POSTGRES_HOST"`
	DatabaseName     string `env:"POSTGRES_DB"`
	DatabaseUser     string `env:"POSTGRES_USER"`
	DatabasePassword string `env:"POSTGRES_PASSWORD"`
	DatabasePort     int    `env:"POSTGRES_CONTAINER_PORT"`
	DatabaseSSLMODE  string `env:"POSTGRES_SSLMODE"`
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

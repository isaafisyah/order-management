package config

import (
	"os"
)

func Get() *Config {
	return &Config{
		Server{
			Name:      os.Getenv("APP_NAME"),
			Host:      os.Getenv("APP_HOST"),
			Port:      os.Getenv("APP_PORT"),
			SecretKey: os.Getenv("APP_SECRET_KEY"),
		},
		Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}

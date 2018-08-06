package config

import (
	"os"
)

type Config struct {
	HttpPort string   `json:"http_port"`
}

func CreateConfig() *Config {
	config := &Config{}

	config.HttpPort = os.Getenv("HTTP_PORT")

	return config
}
package config

import "os"

type ServerConfig struct {
	Host string
	Port string
}

func LoadServerConfig() *ServerConfig {
	serverHost := os.Getenv("SERVER_HOST")
	if serverHost == "" {
		serverHost = "localhost"
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8000"
	}

	return &ServerConfig{
		Host: serverHost,
		Port: serverPort,
	}
}

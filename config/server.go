package config

import "os"

type Environment string

const EnvProduction Environment = "production"
const EnvDevelopment Environment = "development"
const EnvLocal Environment = "local"

type ServerConfig struct {
	Host string
	Port string
	Env  Environment
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

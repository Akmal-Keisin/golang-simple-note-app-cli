package driver

import (
	"simple-note-app/config"
)

type AppConfig struct {
	Server   config.ServerConfig
	Database config.DatabaseConfig
}

func LoadConfig() (*AppConfig, error) {
	serverConfig := config.LoadServerConfig()
	databaseConfig := config.LoadDatabaseConfig()

	return &AppConfig{
		Server:   *serverConfig,
		Database: *databaseConfig,
	}, nil
}

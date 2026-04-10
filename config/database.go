package config

import (
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  bool
}

func LoadDatabaseConfig() *DatabaseConfig {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		dbUsername = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	dbDatabase := os.Getenv("DB_DATABASE")
	if dbDatabase == "" {
		dbDatabase = "bar"
	}

	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		dbSSLMode = "false"
	}

	dbSSLModeBool, err := strconv.ParseBool(dbSSLMode)
	if err != nil {
		fmt.Println("Failed to parse DB_SSL_MODE to boolean: %w", err)
		dbSSLModeBool = false
	}

	return &DatabaseConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		Database: dbDatabase,
		SSLMode:  dbSSLModeBool,
	}
}

package config

import (
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host            string
	Port            string
	Username        string
	Password        string
	Database        string
	SSLMode         bool
	MaxOpenConns    int
	MaxIdleConss    int
	MaxConnLifetime int
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

	maxOpenConns := os.Getenv("DB_MAX_OPEN_CONNS")
	if maxOpenConns == "" {
		maxOpenConns = "25"
	}

	maxOpenConnsInt, err := strconv.Atoi(maxOpenConns)
	if err != nil {
		fmt.Println("Failed to parse DB_MAX_OPEN_CONNS to integer: %w", err)
		maxOpenConnsInt = 25
	}

	maxIdleCons := os.Getenv("DB_MAX_IDLE_CONNS")
	if maxIdleCons == "" {
		maxIdleCons = "25"
	}

	maxIdleConnsInt, err := strconv.Atoi(maxIdleCons)
	if err != nil {
		fmt.Println("Failed to parse DB_MAX_IDLE_CONNS to integer: %w", err)
		maxIdleConnsInt = 25
	}

	maxConnLifetime := os.Getenv("DB_MAX_CONN_LIFETIME")
	if maxConnLifetime == "" {
		maxConnLifetime = "5"
	}

	maxConnLifetimeInt, err := strconv.Atoi(maxConnLifetime)
	if err != nil {
		fmt.Println("Failed to parse DB_MAX_CONN_LIFETIME to integer: %w", err)
		maxConnLifetimeInt = 5
	}

	return &DatabaseConfig{
		Host:            dbHost,
		Port:            dbPort,
		Username:        dbUsername,
		Password:        dbPassword,
		Database:        dbDatabase,
		SSLMode:         dbSSLModeBool,
		MaxOpenConns:    maxOpenConnsInt,
		MaxIdleConss:    maxIdleConnsInt,
		MaxConnLifetime: maxConnLifetimeInt,
	}
}

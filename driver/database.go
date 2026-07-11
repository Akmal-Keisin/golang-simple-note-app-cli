package driver

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseDriver struct {
	DB *sql.DB
}

func StartDatabaseDriver(ctx context.Context, config *AppConfig) (*DatabaseDriver, error) {
	dsn := getDsn(config)

	fmt.Printf(
		"Connecting with DSN: host=%s port=%s user=%s dbname=%s sslmode=%s\n",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Database,
		config.Database.SSLMode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database pool: %w", err)
	}

	// Test connection with 3-second timeout
	pingCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// Ping the database to ensure it's reachable
	err = db.PingContext(pingCtx)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	// Pool configuration
	db.SetMaxOpenConns(config.Database.MaxOpenConns)
	db.SetMaxIdleConns(config.Database.MaxIdleConss)
	db.SetConnMaxLifetime(time.Duration(config.Database.MaxConnLifetime) * time.Minute)

	return &DatabaseDriver{DB: db}, nil
}

func getDsn(config *AppConfig) string {
	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.Database.Username, config.Database.Password),
		Host:   fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port),
		Path:   config.Database.Database,
	}

	query := dsn.Query()
	query.Set("sslmode", config.Database.SSLMode)
	dsn.RawQuery = query.Encode()

	return dsn.String()
}

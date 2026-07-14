package driver

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"simple-note-app/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseDriver struct {
	DB *sql.DB
}

func StartDatabaseDriver(ctx context.Context, appConfig *AppConfig) (*DatabaseDriver, error) {
	dsn := getDsn(appConfig)

	if appConfig.Server.Env == config.EnvDevelopment || appConfig.Server.Env == config.EnvLocal {
		fmt.Printf(
			"Connecting with DSN: host=%s port=%s user=%s dbname=%s sslmode=%s\n",
			appConfig.Database.Host,
			appConfig.Database.Port,
			appConfig.Database.Username,
			appConfig.Database.Database,
			appConfig.Database.SSLMode,
		)
	}

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
	db.SetMaxOpenConns(appConfig.Database.MaxOpenConns)
	db.SetMaxIdleConns(appConfig.Database.MaxIdleConss)
	db.SetConnMaxLifetime(time.Duration(appConfig.Database.MaxConnLifetime) * time.Minute)

	return &DatabaseDriver{DB: db}, nil
}

func getDsn(appConfig *AppConfig) string {
	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(appConfig.Database.Username, appConfig.Database.Password),
		Host:   fmt.Sprintf("%s:%s", appConfig.Database.Host, appConfig.Database.Port),
		Path:   appConfig.Database.Database,
	}

	query := dsn.Query()
	query.Set("sslmode", appConfig.Database.SSLMode)
	dsn.RawQuery = query.Encode()

	return dsn.String()
}

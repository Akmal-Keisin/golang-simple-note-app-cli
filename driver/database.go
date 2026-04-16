package driver

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseDriver struct {
	DB *sql.DB
}

func StartDatabaseDriver(ctx context.Context, config *AppConfig) (*DatabaseDriver, error) {
	dns := getDsn(config)

	db, err := sql.Open("pgx", dns)
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

	return &DatabaseDriver{DB: db}, nil
}

func getDsn(config *AppConfig) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%t",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Database,
		config.Database.SSLMode,
	)

	return dsn
}

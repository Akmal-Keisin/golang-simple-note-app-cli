package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseDriver struct {
	DB *sql.DB
}

func StartDatabaseDriver(config *AppConfig) (*DatabaseDriver, error) {
	dns := getDns(config)

	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	return &DatabaseDriver{DB: db}, nil
}

func getDns(config *AppConfig) string {
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

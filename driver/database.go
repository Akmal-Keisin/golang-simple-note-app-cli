package driver

import (
	"database/sql"
	"fmt"
)

type DatabaseDriver struct {
	DB *sql.DB
}

var DatabaseDriverInstance *DatabaseDriver

func StartDatabaseDriver(config *AppConfig) error {
	dns := getDns(config)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}

	DatabaseDriverInstance = &DatabaseDriver{DB: db}
	return nil
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

package main

import (
	"context"
	"fmt"
	businesslogic "simple-note-app/business-logic"
	"simple-note-app/driver"
	"simple-note-app/helpers"
	"simple-note-app/repositories"
	"simple-note-app/views"

	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variables: %w", err)
	}

	// Load all configurations
	config, err := driver.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config: %w", err)
		return
	}

	fmt.Printf("DB_HOST=%s\n", config.Database.Host)
	fmt.Printf("DB_PORT=%s\n", config.Database.Port)
	fmt.Printf("DB_USERNAME=%s\n", config.Database.Username)
	fmt.Printf("DB_DATABASE=%s\n", config.Database.Database)
	fmt.Printf("DB_SSL_MODE=%s\n", config.Database.SSLMode)

	// Start database connection pool
	db, err := driver.StartDatabaseDriver(ctx, config)
	if err != nil {
		fmt.Println("Failed to start database: %w", err)
		return
	}

	// Ensure the database connection is closed when the application exits
	defer db.DB.Close()

	// for debug purpose
	var databaseName string
	var schemaName string
	var username string

	err = db.DB.QueryRowContext(ctx, "SELECT current_database(), current_schema(), current_user").Scan(
		&databaseName,
		&schemaName,
		&username,
	)

	if err != nil {
		fmt.Printf("Failed to inspect DB connection: %v\n", err)
		return
	}

	fmt.Printf("Connected database: %s\n", databaseName)
	fmt.Printf("Current schema: %s\n", schemaName)
	fmt.Printf("Current user: %s\n", username)
	rows, err := db.DB.QueryContext(ctx, `
		SELECT table_schema, table_name
		FROM information_schema.tables
		WHERE table_name = 'notes'
	`)
	if err != nil {
		fmt.Printf("Failed to inspect tables: %v\n", err)
		return
	}
	defer rows.Close()

	foundNotesTable := false

	for rows.Next() {
		foundNotesTable = true

		var schema string
		var table string

		err := rows.Scan(&schema, &table)
		if err != nil {
			fmt.Printf("Failed to scan table info: %v\n", err)
			return
		}

		fmt.Printf("Found table: %s.%s\n", schema, table)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error while reading table inspection result: %v\n", err)
		return
	}

	if !foundNotesTable {
		fmt.Println("No notes table found in the currently connected database.")
	}

	// Dependency Injection
	repository := repositories.NewRepository(db)
	businessLogic := businesslogic.NewBusinessLogic(repository)
	noteView := views.NewNoteView(businessLogic, repository)

	// Start the application
	helpers.CallClear()
	noteView.Index(ctx)
}

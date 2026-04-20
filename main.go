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

	// Start database connection pool
	db, err := driver.StartDatabaseDriver(ctx, config)
	if err != nil {
		fmt.Println("Failed to start database: %w", err)
		return
	}

	// Ensure the database connection is closed when the application exits
	defer db.DB.Close()

	// Dependency Injection
	repository := repositories.NewRepository(db)
	businessLogic := businesslogic.NewBusinessLogic(repository)
	noteView := views.NewNoteView(businessLogic, repository)

	// Start the application
	helpers.CallClear()
	noteView.Index(ctx)
}

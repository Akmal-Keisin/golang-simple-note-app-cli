package main

import (
	"fmt"
	businesslogic "simple-note-app/business-logic"
	"simple-note-app/driver"
	"simple-note-app/helpers"
	"simple-note-app/repositories"
	"simple-note-app/views"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variables: %w", err)
	}

	config, err := driver.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config: %w", err)
		return
	}

	db, err := driver.StartDatabaseDriver(config)
	if err != nil {
		fmt.Println("Failed to start database: %w", err)
		return
	}

	helpers.CallClear()

	// Dependency Injection
	repository := repositories.NewRepository(db)
	businessLogic := businesslogic.NewBusinessLogic(repository)
	noteView := views.NewNoteView(businessLogic)
	noteView.Index()
}

package main

import (
	"fmt"
	"simple-note-app/driver"
	"simple-note-app/helpers"
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

	err = driver.StartDatabaseDriver(config)
	if err != nil {
		fmt.Println("Failed to start database: %w", err)
		return
	}

	helpers.CallClear()
	views.Index()
}

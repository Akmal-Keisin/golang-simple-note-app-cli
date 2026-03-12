package main

import (
	"fmt"
	"simple-note-app/driver"
	"simple-note-app/helpers"
	"simple-note-app/views"
)

func main() {
	_, err := driver.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config: %w", err)
		return
	}

	helpers.CallClear()
	views.Index()
}

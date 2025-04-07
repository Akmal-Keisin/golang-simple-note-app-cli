package views

import (
	"bufio"
	"fmt"
	"os"
	"simple-note-app/helpers"
	"simple-note-app/model"
	"strings"
)

func Read() {
	helpers.CallClear()
	fmt.Println("Read Note")
	for i, note := range model.Notes {
		fmt.Printf("%s %d\n", "Note", i+1)
		fmt.Printf("ID: %d\n", note.Id)
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)
		fmt.Printf("Created At: %s\n", note.CreatedAt)
		fmt.Printf("Updated At: %s\n", note.UpdatedAt)
		fmt.Println("===================================")
	}

	confirmBackReader := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to go back to menu... ")
	confirm, _ := confirmBackReader.ReadString('\n')
	
	if strings.TrimSpace(confirm) == "" {
		Index()
	} else {
		fmt.Println("Invalid input, please try again.")
		Read()
	}
}
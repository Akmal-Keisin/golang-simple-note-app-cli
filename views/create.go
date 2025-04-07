package views

import (
	"bufio"
	"fmt"
	"os"
	businesslogic "simple-note-app/business-logic"
	"simple-note-app/helpers"
	"strings"
)

func Create() {
	helpers.CallClear()
	fmt.Println("Create New Note")

	titleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Note Title: ")
	title, _ := titleReader.ReadString('\n')
	title = strings.TrimSpace(title)

	contentReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Note Content: ")
	content, _ := contentReader.ReadString('\n')
	content = strings.TrimSpace(content)

	fmt.Println("You are trying to input note with")
	fmt.Print("Title: ", title)
	fmt.Print("Content: ", content)
	fmt.Println("Are you sure? (y/n)")

	confirmReader := bufio.NewReader(os.Stdin)
	fmt.Print("Confirm: ")
	confirm, _ := confirmReader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm != "y" {
		fmt.Println("Note creation cancelled.")
		Index("Note creation cancelled.")
	}

	businesslogic.HandleCreateNote(title, content)
	Index("Note Created Successfully!")
}
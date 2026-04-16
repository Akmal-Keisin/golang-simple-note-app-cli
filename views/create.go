package views

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"simple-note-app/helpers"
	"strings"
)

func (n *NoteView) Create(ctx context.Context, flashMessage ...string) {
	helpers.CallClear()
	fmt.Println("Create New Note")

	fmt.Print("Enter Note Title: ")
	titleReader := bufio.NewReader(os.Stdin)
	title, _ := titleReader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter Note Content: ")
	contentReader := bufio.NewReader(os.Stdin)
	content, _ := contentReader.ReadString('\n')
	content = strings.TrimSpace(content)

	fmt.Println("You are trying to input note with")
	fmt.Print("Title: ", title)
	fmt.Print("Content: ", content)
	fmt.Println("Are you sure? (y/n)")

	fmt.Print("Confirm: ")
	confirmReader := bufio.NewReader(os.Stdin)
	confirm, _ := confirmReader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm != "y" {
		fmt.Println("Note creation cancelled.")
		n.Index(ctx, "Note creation cancelled.")
	}

	n.businessLogic.HandleCreateNote(ctx, title, content)
	n.Index(ctx, "Note Created Successfully!")
}

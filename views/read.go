package views

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"simple-note-app/helpers"
	"strings"
)

func (n *NoteView) Read(ctx context.Context, flashMessage ...string) {
	helpers.CallClear()
	fmt.Println("Read Note")

	// Show all notes to let the user select a note to see the detail
	notes, err := n.repository.GetAllNotes(ctx)
	if err != nil {
		fmt.Printf("An error occured: %v \n", err)
		return
	}

	for _, note := range notes {
		fmt.Printf("ID: %d\n", note.Id)
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)

		if note.CreatedAt != nil {
			fmt.Printf("Created At: %s\n", *note.CreatedAt)
		} else {
			fmt.Println("Created At: -")
		}

		if note.UpdatedAt != nil {
			fmt.Printf("Updated At: %s\n", *note.UpdatedAt)
		} else {
			fmt.Println("Updated At: -")
		}

		fmt.Println("===================================")
	}

	confirmBackReader := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to go back to menu... ")
	confirm, _ := confirmBackReader.ReadString('\n')

	if strings.TrimSpace(confirm) == "" {
		n.Index(ctx)
	} else {
		fmt.Println("Invalid input, please try again.")
		n.Read(ctx)
	}
}

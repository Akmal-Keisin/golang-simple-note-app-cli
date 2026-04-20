package views

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"simple-note-app/helpers"
	"strconv"
	"strings"
)

func (n *NoteView) Edit(ctx context.Context, flashMessage ...string) {
	helpers.CallClear()
	fmt.Println("Edit Note Note")

	// Show all notes to let the user select a note to update
	notes, err := n.repository.GetAllNotes(ctx)
	if err != nil {
		fmt.Println("An error occured")
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

	selectNoteReader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose Note ID To Edit: ")
	noteId, _ := selectNoteReader.ReadString('\n')
	noteId = strings.TrimSpace(noteId)
	noteIdInt, err := strconv.Atoi(noteId)

	if err != nil {
		fmt.Println("Invalid Note ID")
		n.Edit(ctx)
	}

	var newTitle string
	var newContent string

	// Provide the user update input
	for _, note := range notes {
		if note.Id == noteIdInt {
			helpers.CallClear()
			fmt.Printf("You choosed Note ID %d\n", noteIdInt)
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
			fmt.Println("Edit Note")

			editTitle := bufio.NewReader(os.Stdin)
			fmt.Print("New Title: ")
			selectedNewTitle, _ := editTitle.ReadString('\n')
			selectedNewTitle = strings.TrimSpace(selectedNewTitle)

			editContent := bufio.NewReader(os.Stdin)
			fmt.Print("New Content: ")
			selectedNewContent, _ := editContent.ReadString('\n')
			selectedNewContent = strings.TrimSpace(selectedNewContent)

			noteIdInt = note.Id
			newTitle = selectedNewTitle
			newContent = selectedNewContent
			break
		}
	}

	confirmUpdate := bufio.NewReader(os.Stdin)
	fmt.Print("You are about to update this note, are you sure? (y/n): ")
	confirm, _ := confirmUpdate.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm != "y" {
		fmt.Println("Note update cancelled.")
		n.Edit(ctx)
	}

	msg, err := n.businessLogic.HandleUpdateNote(ctx, noteIdInt, newTitle, newContent)
	if err != nil {
		fmt.Printf("An error occured: %v \n", err)
		return
	}
	n.Index(ctx, msg)
}

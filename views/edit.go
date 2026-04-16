package views

import (
	"bufio"
	"fmt"
	"os"
	"simple-note-app/helpers"
	"simple-note-app/model"
	"strconv"
	"strings"
	"time"
)

func (n *NoteView) Edit(flashMessage ...string) {
	helpers.CallClear()
	fmt.Println("Edit Note Note")
	for i, note := range model.Notes {
		fmt.Printf("%s %d\n", "Note", i+1)
		fmt.Printf("ID: %d\n", note.Id)
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)
		fmt.Printf("Created At: %s\n", note.CreatedAt)
		fmt.Println("===================================")
	}

	selectNoteReader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose Note ID To Edit: ")
	noteId, _ := selectNoteReader.ReadString('\n')
	noteId = strings.TrimSpace(noteId)
	noteIdInt, err := strconv.Atoi(noteId)

	if err != nil {
		fmt.Println("Invalid Note ID")
		n.Edit()
	}

	updatedNote := model.Note{}

	for _, note := range model.Notes {
		if note.Id == noteIdInt {
			helpers.CallClear()
			fmt.Printf("You choosed Note ID %d\n", noteIdInt)
			fmt.Printf("Title: %s\n", note.Title)
			fmt.Printf("Content: %s\n", note.Content)
			fmt.Printf("Created At: %s\n", note.CreatedAt)
			fmt.Printf("Updated At: %s\n", note.UpdatedAt)
			fmt.Println("===================================")
			fmt.Println("Edit Note")

			editTitle := bufio.NewReader(os.Stdin)
			fmt.Print("New Title: ")
			newTitle, _ := editTitle.ReadString('\n')
			newTitle = strings.TrimSpace(newTitle)

			editContent := bufio.NewReader(os.Stdin)
			fmt.Print("New Content: ")
			newContent, _ := editContent.ReadString('\n')
			newContent = strings.TrimSpace(newContent)

			updatedNote.Id = note.Id
			updatedNote.Title = newTitle
			updatedNote.Content = newContent
			updatedNote.CreatedAt = note.CreatedAt
			updatedNote.UpdatedAt = time.Now().Format(time.RFC850)
			break
		}
	}

	confirmUpdate := bufio.NewReader(os.Stdin)
	fmt.Print("You are about to update this note, are you sure? (y/n): ")
	confirm, _ := confirmUpdate.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm != "y" {
		fmt.Println("Note update cancelled.")
		n.Edit()
	}

	n.businessLogic.HandleUpdateNote(noteIdInt, updatedNote)
	n.Index("Note Updated Successfully!")
}

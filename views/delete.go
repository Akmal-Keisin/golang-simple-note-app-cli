package views

import (
	"bufio"
	"fmt"
	"os"
	"simple-note-app/helpers"
	"simple-note-app/model"
	"strconv"
	"strings"
)

func (n *NoteView) Delete(flashMessage ...string) {
	helpers.CallClear()
	fmt.Println("Delete Note")
	for i, note := range model.Notes {
		fmt.Printf("%s %d\n", "Note", i+1)
		fmt.Printf("ID: %d\n", note.Id)
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)
		fmt.Printf("Created At: %s\n", note.CreatedAt)
		fmt.Printf("Updated At: %s\n", note.UpdatedAt)
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

	findNote := model.Note{}
	for _, note := range model.Notes {
		if note.Id == noteIdInt {
			helpers.CallClear()
			fmt.Printf("You choosed Note ID %d\n", noteIdInt)
			fmt.Printf("Title: %s\n", note.Title)
			fmt.Printf("Content: %s\n", note.Content)
			fmt.Printf("Created At: %s\n", note.CreatedAt)
			fmt.Printf("Updated At: %s\n", note.UpdatedAt)
			fmt.Println("===================================")
			fmt.Println("Delete Note")
			findNote = note
			break
		}
	}

	if findNote.Id == 0 {
		fmt.Println("Note not found")
		n.Delete()
	}

	confirmDelete := bufio.NewReader(os.Stdin)
	fmt.Print("You are about to delete this note, are you sure? (y/n): ")
	confirm, _ := confirmDelete.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm != "y" {
		fmt.Println("Note delete cancelled.")
		n.Edit()
	}

	n.businessLogic.HandleDeleteNote(findNote.Id)
	n.Index("Note Deleted Successfully!")
}

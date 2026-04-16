package businesslogic

import (
	"simple-note-app/repositories"
	"time"
)

func (businessLogic *BusinessLogic) HandleCreateNote(title string, content string) (message string, err error) {
	timestamp := time.Now().Format(time.RFC850)

	newNoteRequest := repositories.CreateNoteRequest{
		Title:     title,
		Content:   content,
		CreatedAt: timestamp,
	}

	newNote, err := businessLogic.repository.CreateNote(newNoteRequest)
	if err != nil {
		return "", err
	}

	return "Note created successfully with ID: " + string(newNote.Id), nil
}

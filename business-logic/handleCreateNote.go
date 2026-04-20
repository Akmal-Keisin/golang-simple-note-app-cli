package businesslogic

import (
	"context"
	"fmt"
	"simple-note-app/repositories"
	"time"
)

func (businessLogic *BusinessLogic) HandleCreateNote(ctx context.Context, title string, content string) (message string, err error) {
	currentDate := time.Now().Format(time.DateOnly)
	currentTime := time.Now().Format(time.TimeOnly)

	newNoteRequest := repositories.CreateNoteRequest{
		Title:     title,
		Content:   content,
		CreatedAt: fmt.Sprintf("%s %s", currentDate, currentTime),
	}

	newNote, err := businessLogic.repository.CreateNote(ctx, newNoteRequest)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Note created successfully with ID: %d", newNote.Id), nil
}

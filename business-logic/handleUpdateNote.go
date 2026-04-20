package businesslogic

import (
	"context"
	"fmt"
	"simple-note-app/repositories"
	"time"
)

func (businessLogic *BusinessLogic) HandleUpdateNote(ctx context.Context, noteId int, title string, content string) (message string, err error) {
	currentDate := time.Now().Format(time.DateOnly)
	currentTime := time.Now().Format(time.TimeOnly)

	updateNoteRequest := repositories.UpdateNoteRequest{
		Id:        noteId,
		Title:     title,
		Content:   content,
		UpdatedAt: fmt.Sprintf("%s %s", currentDate, currentTime),
	}

	updateNote, err := businessLogic.repository.UpdateNote(ctx, updateNoteRequest)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Note updated successfully with ID: %d", updateNote.Id), nil

}

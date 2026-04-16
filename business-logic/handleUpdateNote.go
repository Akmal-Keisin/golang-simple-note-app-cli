package businesslogic

import (
	"context"
	"simple-note-app/model"
)

func (businessLogic *BusinessLogic) HandleUpdateNote(ctx context.Context, noteId int, updatedNote model.Note) (message string, err error) {

	// TODO: Validate if the note exists

	// TODO: Update the note using repository

	for i, note := range model.Notes {
		if note.Id == noteId {
			model.Notes[i] = updatedNote
			break
		}
	}

	return "", nil
}

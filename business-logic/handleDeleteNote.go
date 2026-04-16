package businesslogic

import (
	"context"
	"simple-note-app/model"
)

func (businessLogic *BusinessLogic) HandleDeleteNote(ctx context.Context, noteId int) (message string, err error) {
	// TODO: Validate if the note exists

	// TODO: Delete the note using repository

	indexToDelete := -1
	for i, note := range model.Notes {
		if note.Id == noteId {
			indexToDelete = i
			break
		}
	}

	if indexToDelete != -1 {
		model.Notes = append(model.Notes[:indexToDelete], model.Notes[indexToDelete+1:]...)
	}

	return "", nil
}

package businesslogic

import (
	"simple-note-app/model"
)

func (businessLogic *BusinessLogic) HandleUpdateNote(noteId int, updatedNote model.Note) (message string, err error) {

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

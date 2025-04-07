package businesslogic

import (
	"simple-note-app/model"
)

func HandleUpdateNote(noteId int, updatedNote model.Note) {
	for i, note := range model.Notes {
		if note.Id == noteId {
			model.Notes[i] = updatedNote
			break;
		}
	}
}
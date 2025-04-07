package businesslogic

import "simple-note-app/model"

func HandleDeleteNote(noteId int) {
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
}
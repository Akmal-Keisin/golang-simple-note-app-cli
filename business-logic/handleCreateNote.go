package businesslogic

import (
	"simple-note-app/model"
	"time"
)

func HandleCreateNote(title string, content string) {
	timestamp := time.Now().Format(time.RFC850)

	newNoteId := 1;
	lastNote := model.Notes[len(model.Notes)-1]
	if lastNote.Id != 0 {
		newNoteId = lastNote.Id + 1
	}
	
	newNote := model.Note{
		Id: newNoteId,
		Title:     title,
		Content:   content,
		CreatedAt: timestamp,
	}

	model.Notes = append(model.Notes, newNote)
}
package repositories

import "simple-note-app/model"

type CreateNoteRequest struct {
	Title     string
	Content   string
	CreatedAt string
}

func (repository *Repository) CreateNote(request CreateNoteRequest) (*model.Note, error) {
	// TODO: implement new note query
	return nil, nil
}

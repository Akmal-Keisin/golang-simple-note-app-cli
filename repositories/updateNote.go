package repositories

import "simple-note-app/model"

type UpdateNoteRequest struct {
	Id        int
	Title     string
	Content   string
	UpdatedAt string
}

func (repository *Repository) UpdateNote(request UpdateNoteRequest) (*model.Note, error) {
	// TODO: implement update note query
	return nil, nil
}

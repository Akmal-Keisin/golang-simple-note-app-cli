package repositories

import "simple-note-app/model"

type NoteRepositoryInterface interface {
	Get() []model.Note
	Find() model.Note
	Create() int
	Update() int
	Delete() int
}

type NoteRepository struct {
}

func NewNoteRepository() NoteRepositoryInterface {
	return &NoteRepository{}
}

func (noteRepository NoteRepository) Get() []model.Note {
	return nil
}

func (noteRepository NoteRepository) Find() model.Note {
	return model.Note{}
}

func (noteRepository NoteRepository) Create() int {
	return 0
}

func (noteRepository NoteRepository) Update() int {
	return 0
}

func (noteRepository NoteRepository) Delete() int {
	return 0
}

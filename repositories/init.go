package repositories

import (
	"simple-note-app/driver"
	"simple-note-app/model"
)

type RepositoryInterface interface {
	GetAllNotes() ([]model.Note, error)
	FindNoteByID(id int) (*model.Note, error)
	CreateNote(CreateNoteRequest) (*model.Note, error)
	UpdateNote(UpdateNoteRequest) (*model.Note, error)
	DeleteNote(id int) error
}

type Repository struct {
	db *driver.DatabaseDriver
}

func NewRepository(db *driver.DatabaseDriver) RepositoryInterface {
	return &Repository{db: db}
}

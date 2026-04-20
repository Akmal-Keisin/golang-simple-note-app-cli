package repositories

import (
	"context"
	"simple-note-app/driver"
	"simple-note-app/model"
)

type RepositoryInterface interface {
	GetAllNotes(ctx context.Context) ([]model.Note, error)
	FindNoteByID(ctx context.Context, id int) (*model.Note, error)
	CreateNote(ctx context.Context, request CreateNoteRequest) (*model.Note, error)
	UpdateNote(ctx context.Context, request UpdateNoteRequest) (*model.Note, error)
	DeleteNote(ctx context.Context, id int) (*model.Note, error)
}

type Repository struct {
	db *driver.DatabaseDriver
}

func NewRepository(db *driver.DatabaseDriver) RepositoryInterface {
	return &Repository{db: db}
}

package repositories

import (
	"context"
	"fmt"
	"simple-note-app/model"
)

type CreateNoteRequest struct {
	Title     string
	Content   string
	CreatedAt string
}

func (repository *Repository) CreateNote(ctx context.Context, request CreateNoteRequest) (*model.Note, error) {

	query := "INSERT INTO notes (title, content, created_at) VALUES ($1, $2, $3) RETURNING id, title, content, created_at"

	row := repository.db.DB.QueryRowContext(ctx, query, request.Title, request.Content, request.CreatedAt)
	if row.Err() != nil {
		return nil, fmt.Errorf("Failed to create a new note : %v", row.Err())
	}

	newNote := model.Note{}
	row.Scan(&newNote.Id, &newNote.Title, &newNote.Content, &newNote.CreatedAt)

	return &newNote, nil
}

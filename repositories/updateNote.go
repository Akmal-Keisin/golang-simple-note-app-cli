package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"simple-note-app/model"
)

type UpdateNoteRequest struct {
	Id        int
	Title     string
	Content   string
	UpdatedAt string
}

func (repository *Repository) UpdateNote(ctx context.Context, request UpdateNoteRequest) (*model.Note, error) {

	query := "UPDATE notes SET title = $1, content = $2, updated_at = $3 WHERE id = $4 RETURNING id, title, content, created_at, updated_at"

	row := repository.db.DB.QueryRowContext(ctx, query, request.Title, request.Content, request.UpdatedAt, request.Id)
	err := row.Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Failed to update note : %v", err)
		}

		return nil, err
	}

	updatedNote := model.Note{}
	row.Scan(
		&updatedNote.Id,
		&updatedNote.Title,
		&updatedNote.Content,
		&updatedNote.CreatedAt,
		&updatedNote.UpdatedAt,
	)

	return &updatedNote, nil
}

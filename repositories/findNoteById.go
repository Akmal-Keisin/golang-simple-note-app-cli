package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"simple-note-app/model"
)

func (repository *Repository) FindNoteByID(ctx context.Context, id int) (*model.Note, error) {

	query := "SELECT id, title, content, created_at, udpated_at FROM notes WHERE id = $1"

	row := repository.db.DB.QueryRowContext(ctx, query, id)
	err := row.Err()

	if err != nil {

		// Not found error
		if err == sql.ErrNoRows {
			return nil, nil
		}

		// Unexpected error
		return nil, fmt.Errorf("Failed to find note by id : %v", err)
	}

	note := model.Note{}
	row.Scan(
		&note.Id,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	return &note, nil
}

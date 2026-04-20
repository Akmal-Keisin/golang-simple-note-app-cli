package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"simple-note-app/model"
)

func (repository *Repository) DeleteNote(ctx context.Context, id int) (*model.Note, error) {

	query := "DELETE FROM notes WHERE id = $1 RETURNING id, title, content, created_at, updated_at"

	row := repository.db.DB.QueryRowContext(ctx, query, id)
	err := row.Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Failed to delete note : %v", err)
		}

		return nil, err
	}

	deletedNote := model.Note{}
	row.Scan(
		&deletedNote.Id,
		&deletedNote.Title,
		&deletedNote.Content,
		&deletedNote.CreatedAt,
		&deletedNote.UpdatedAt,
	)

	return &deletedNote, nil
}

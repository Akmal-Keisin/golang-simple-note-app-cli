package repositories

import (
	"context"
	"fmt"
	"simple-note-app/model"
)

func (repository *Repository) GetAllNotes(ctx context.Context) ([]model.Note, error) {

	query := "SELECT id, title, content, created_at, updated_at FROM notes ORDER BY title"

	rows, err := repository.db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get all notes : %v", err)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Failed to get all notes : %v", err)
	}

	defer rows.Close()

	notes := []model.Note{}
	for rows.Next() {
		note := model.Note{}
		if err := rows.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Failed to scan rows : %v", err)
		}

		notes = append(notes, note)
	}

	return notes, nil
}

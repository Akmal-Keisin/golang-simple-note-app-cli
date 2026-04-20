package businesslogic

import (
	"context"
	"fmt"
)

func (businessLogic *BusinessLogic) HandleDeleteNote(ctx context.Context, noteId int) (message string, err error) {
	deletedNote, err := businessLogic.repository.DeleteNote(ctx, noteId)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Note deleted successfully with ID: %d", deletedNote.Id), nil
}

package businesslogic

import (
	"context"
	"simple-note-app/repositories"
)

type BusinessLogicInterface interface {
	HandleCreateNote(ctx context.Context, title string, content string) (message string, err error)
	HandleUpdateNote(ctx context.Context, noteId int, title string, content string) (message string, err error)
	HandleDeleteNote(ctx context.Context, noteId int) (message string, err error)
}

type BusinessLogic struct {
	repository repositories.RepositoryInterface
}

func NewBusinessLogic(repository repositories.RepositoryInterface) BusinessLogicInterface {
	return &BusinessLogic{repository: repository}
}

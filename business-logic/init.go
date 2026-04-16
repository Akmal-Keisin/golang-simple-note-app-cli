package businesslogic

import (
	"simple-note-app/model"
	"simple-note-app/repositories"
)

type BusinessLogicInterface interface {
	HandleCreateNote(title string, content string) (message string, err error)
	HandleUpdateNote(noteId int, updatedNote model.Note) (message string, err error)
	HandleDeleteNote(noteId int) (message string, err error)
}

type BusinessLogic struct {
	repository repositories.RepositoryInterface
}

func NewBusinessLogic(repository repositories.RepositoryInterface) BusinessLogicInterface {
	return &BusinessLogic{repository: repository}
}

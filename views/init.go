package views

import (
	"context"
	businesslogic "simple-note-app/business-logic"
	"simple-note-app/repositories"
)

type NoteViewInterface interface {
	Index(ctx context.Context, flashMessage ...string)
	Create(ctx context.Context, flashMessage ...string)
	Read(ctx context.Context, flashMessage ...string)
	Edit(ctx context.Context, flashMessage ...string)
	Delete(ctx context.Context, flashMessage ...string)
}

type NoteView struct {
	businessLogic businesslogic.BusinessLogicInterface
	repository    repositories.RepositoryInterface
}

func NewNoteView(businessLogic businesslogic.BusinessLogicInterface, repository repositories.RepositoryInterface) NoteViewInterface {
	return &NoteView{
		businessLogic: businessLogic,
		repository:    repository,
	}
}

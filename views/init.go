package views

import (
	"context"
	businesslogic "simple-note-app/business-logic"
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
}

func NewNoteView(businessLogic businesslogic.BusinessLogicInterface) NoteViewInterface {
	return &NoteView{
		businessLogic: businessLogic,
	}
}

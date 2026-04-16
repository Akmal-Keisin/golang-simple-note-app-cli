package views

import businesslogic "simple-note-app/business-logic"

type NoteViewInterface interface {
	Index(flashMessage ...string)
	Create(flashMessage ...string)
	Read(flashMessage ...string)
	Edit(flashMessage ...string)
	Delete(flashMessage ...string)
}

type NoteView struct {
	businessLogic businesslogic.BusinessLogicInterface
}

func NewNoteView(businessLogic businesslogic.BusinessLogicInterface) NoteViewInterface {
	return &NoteView{
		businessLogic: businessLogic,
	}
}

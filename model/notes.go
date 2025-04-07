package model

import "time"

type Note struct {
	Id        int
	Title     string
	Content   string
	CreatedAt string
	UpdatedAt string
}

var Notes []Note = []Note{}

func init() {
	Notes = append(
		Notes, Note{
			Id: 1, 
			Title: "First Note", 
			Content: "This is the content of the first note", 
			CreatedAt: time.Now().Format(time.RFC850),
		},
	)

	Notes = append(
		Notes, Note{
			Id: 2, 
			Title: "Second Note", 
			Content: "This is the content of the second note", 
			CreatedAt: time.Now().Format(time.RFC850),
		},
	)

	Notes = append(
		Notes, Note{
			Id: 3, 
			Title: "Third Note", 
			Content: "This is the content of the third note", 
			CreatedAt: time.Now().Format(time.RFC850),
		},
	)
}
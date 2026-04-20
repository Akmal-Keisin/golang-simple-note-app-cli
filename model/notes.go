package model

import "time"

type Note struct {
	Id        int
	Title     string
	Content   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

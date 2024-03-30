package domain

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Text     string
	Status   Status
	Deadline int
}

type Status int

const (
	Task Status = iota
	ThisWeek
	Doing
	Review
	Done
	Close
)

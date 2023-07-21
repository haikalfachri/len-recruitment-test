package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string      `json:"name"`
	SID     string      `json:"sid"`
	Major   string      `json:"major"`
	Borrows []Borrowing `json:"-" gorm:"foreignKey:StudentID"`
}

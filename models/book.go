package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string      `json:"title"`
	Author   string      `json:"author"`
	Quantity int64       `json:"quantity"`
	Storage  string      `json:"storage"`
	Borrows  []Borrowing `json:"-" gorm:"foreignKey:BookID"`
}

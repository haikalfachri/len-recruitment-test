package models

import (
	"time"

	"gorm.io/gorm"
)

type Borrowing struct {
	gorm.Model
	StudentID    uint      `json:"student_id"`
	BookID       uint      `json:"book_id"`
	BorrowedAt   time.Time `json:"borrowed_at"`
	BorrowedTerm time.Time `json:"borrowed_term"`
	ReturnedAt   time.Time `json:"returned_at"`
}

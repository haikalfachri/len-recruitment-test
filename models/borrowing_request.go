package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type BorrowingRequest struct {
	StudentID    uint      `json:"student_id" validate:"required"`
	BookID       uint      `json:"book_id" validate:"required"`
	BorrowedTerm time.Time `json:"borrowed_term" validate:"required"`
	ReturnedAt   time.Time `json:"returned_at" gorm:"default:CURRENT_TIMESTAMP()"`
}

type ReturningRequest struct {
	ReturnedAt   time.Time `json:"returned_at" validate:"required"`
}

func (u *BorrowingRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

func (u *ReturningRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

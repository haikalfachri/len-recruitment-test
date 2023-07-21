package models

import "github.com/go-playground/validator/v10"

type BookRequest struct {
	Title    string `json:"title" validate:"required"`
	Author   string `json:"author" validate:"required"`
	Quantity int64 `json:"quantity" validate:"required"`
	Storage  string `json:"storage" validate:"required"`
}

func (u *BookRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}
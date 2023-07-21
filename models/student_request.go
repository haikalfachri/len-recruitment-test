package models

import "github.com/go-playground/validator/v10"

type StudentRequest struct {
	Name  string `json:"name" validate:"required"`
	SID   string `json:"sid" validate:"required"`
	Major string `json:"major" validate:"required"`
}

func (u *StudentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

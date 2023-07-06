package models

import (
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required,min=20"`
	Slug        string `json:"slug"`
	Content     string `json:"content" validate:"required,min=200"`
	Category    string `json:"category" validate:"required,min=3"`
	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
	Status      string `json:"status" validate:"required,oneof=Publish Draft Thrash"`
}

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Properti dibutuhkan."
	case "min":
		return "Minimal karakter adalah " + fe.Param()
	}
	return fe.Error() // default error
}

package model

type User struct {
	Model[string]

	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`

	IsAdmin bool `json:"is_admin"`

	Active bool `json:"active"`
}

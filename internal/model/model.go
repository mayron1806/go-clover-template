package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Model[K comparable] struct {
	ID        K         `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Model[K]) Validate(opts ...validator.Option) error {
	validate := validator.New(opts...)
	return validate.Struct(m)
}
func (m *Model[K]) Update(new *Model[K]) error {
	id := m.ID
	m = new
	m.ID = id

	return m.Validate()
}

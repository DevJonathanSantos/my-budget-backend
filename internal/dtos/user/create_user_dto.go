package dto

import (
	"errors"
	"net/mail"
	"time"
)

type CreateUserRequestDto struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type CreateUserResponseDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func (dto *CreateUserRequestDto) Validate() error {
	if dto.Name == "" {
		return errors.New("name is required")
	}
	if len(dto.Name) < 3 {
		return errors.New("name must be at least 3 characters long")
	}
	if dto.Email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(dto.Email); err != nil {
		return errors.New("invalid email format")
	}

	return nil
}

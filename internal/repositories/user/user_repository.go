package repositories

import (
	"context"
	"time"

	"github.com/DevJonathanSantos/my-budget-backend/internal/entities"
	"github.com/google/uuid"
)

func (r *userRepository) Create(ctx context.Context, user *entities.UserEntity) (*entities.UserEntity, error) {
	userCreated := entities.UserEntity{
		ID:        uuid.New().String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: time.Now(),
	}

	return &userCreated, nil
}

package user_use_cases

import (
	"context"

	dto "github.com/DevJonathanSantos/my-budget-backend/internal/dtos/user"
	repository "github.com/DevJonathanSantos/my-budget-backend/internal/repositories/user"
)

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo,
	}
}

type userUseCase struct {
	repo repository.UserRepository
}

type UserUseCase interface {
	Create(ctx context.Context, request *dto.CreateUserRequestDto) (*dto.CreateUserResponseDto, error)
}

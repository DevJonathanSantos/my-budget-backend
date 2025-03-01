package user_use_cases

import (
	"context"
	"fmt"

	dto "github.com/DevJonathanSantos/my-budget-backend/internal/dtos/user"
	"github.com/DevJonathanSantos/my-budget-backend/internal/entities"
)

func (u *userUseCase) Create(ctx context.Context, request *dto.CreateUserRequestDto) (*dto.CreateUserResponseDto, error) {

	user := entities.UserEntity{
		Name:  request.Name,
		Email: request.Email,
	}

	userCreated, _ := u.repo.Create(ctx, &user)

	fmt.Println(userCreated)

	fmt.Printf("user da requisição: %s\n", user)
	fmt.Printf("Resposta da requisição: %s\n", *userCreated)

	res := dto.CreateUserResponseDto{
		ID:    userCreated.ID,
		Name:  userCreated.Name,
		Email: userCreated.Email,
	}

	return &res, nil
}

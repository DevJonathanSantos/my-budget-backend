package handlers

import (
	"context"

	use_cases "github.com/DevJonathanSantos/my-budget-backend/internal/use_cases/user"
	"github.com/aws/aws-lambda-go/events"
)

func NewUserHandler(useCase use_cases.UserUseCase) UserHandler {
	return &userHandler{
		useCase: useCase,
	}
}

type userHandler struct {
	useCase use_cases.UserUseCase
}

type UserHandler interface {
	CreateUserHandler(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

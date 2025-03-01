package main

import (
	"github.com/DevJonathanSantos/my-budget-backend/internal/database"

	handlers "github.com/DevJonathanSantos/my-budget-backend/internal/handlers/user"
	user_repository "github.com/DevJonathanSantos/my-budget-backend/internal/repositories/user"
	use_cases "github.com/DevJonathanSantos/my-budget-backend/internal/use_cases/user"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	db := database.NewDBConnection()
	repository := user_repository.NewUserRepository(db, "users-db")
	useCase := use_cases.NewUserUseCase(repository)
	handler := handlers.NewUserHandler(useCase)
	lambda.Start(handler.CreateUserHandler)
}

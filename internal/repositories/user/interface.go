package repositories

import (
	"context"

	"github.com/DevJonathanSantos/my-budget-backend/internal/entities"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewUserRepository(db *dynamodb.Client, tableName string) UserRepository {
	return &userRepository{
		db:        db,
		tableName: tableName,
	}
}

type userRepository struct {
	db        *dynamodb.Client
	tableName string
}

type UserRepository interface {
	Create(ctx context.Context, user *entities.UserEntity) (*entities.UserEntity, error)
}

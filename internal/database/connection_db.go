package database

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	dynamoClient *dynamodb.Client
	once         sync.Once
)

func NewDBConnection() *dynamodb.Client {
	once.Do(func() {
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion("us-east-1"),
		)
		if err != nil {
			log.Fatalf("Erro ao carregar configuração da AWS: %v", err)
		}

		dynamoClient = dynamodb.NewFromConfig(cfg)
	})

	return dynamoClient
}

package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	dto "github.com/DevJonathanSantos/my-budget-backend/internal/dtos/user"
	"github.com/aws/aws-lambda-go/events"
)

const (
	StatusCodeOK            = 200
	StatusCodeCreated       = 201
	StatusCodeBadRequest    = 400
	StatusCodeInternalError = 500
)

func (h *userHandler) CreateUserHandler(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Recebendo requisição para criação de usuário")
	fmt.Printf("Corpo da requisição: %s\n", request.Body)

	var requestDto dto.CreateUserRequestDto
	if err := json.Unmarshal([]byte(request.Body), &requestDto); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return jsonResponse(StatusCodeBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := requestDto.Validate(); err != nil {
		fmt.Println("Erro de validação:", err)
		return jsonResponse(StatusCodeBadRequest, map[string]string{"error": err.Error()})
	}

	responseDto, err := h.useCase.Create(ctx, &requestDto)
	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		return jsonResponse(StatusCodeInternalError, map[string]string{"error": "Failed to create user"})
	}

	fmt.Println("Usuário criado com sucesso!")
	return jsonResponse(StatusCodeCreated, responseDto)
}

func jsonResponse(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	responseBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Erro ao processar resposta JSON:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: StatusCodeInternalError,
			Body:       `{"error": "Failed to process response"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(responseBody),
	}, nil
}

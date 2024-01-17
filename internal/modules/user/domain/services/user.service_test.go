package services_test

import (
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/mocks"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/golang/mock/gomock"
)

func prepare(t *testing.T) (*services.UserService, *mocks.MockUserRepository, *mocks.MockIRabbitMQ) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockUserRepository(ctrl)
	mockRabbitMQ := mocks.NewMockIRabbitMQ(ctrl)
	userService := services.NewUserService(mockUserRepository, mockRabbitMQ)

	return userService, mockUserRepository, mockRabbitMQ
}

var id, err = pkg_entities.ParseID("52c599f3-af83-4fd9-bfd6-e532918f7b13")
var createdAt, _ = time.Parse(time.RFC3339Nano, "2024-01-17T01:04:23.883938Z")
var password = "hgGFHJ654*"
var user = &entities.User{
	ID:        id,
	Name:      "John",
	LastName:  "Doe",
	CPF:       "123.356.229-00",
	Email:     "john.doe@example.com",
	Password:  password,
	Admin:     false,
	CreatedAt: createdAt,
	UpdatedAt: nil,
}
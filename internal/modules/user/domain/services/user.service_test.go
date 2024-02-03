package services_test

import (
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/mocks"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/golang/mock/gomock"
	"github.com/streadway/amqp"
)

func prepare(t testing.TB) (*services.UserService, *mocks.MockUserRepository, *mocks.MockIRabbitMQ, *mocks.MockMencacher) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockUserRepository(ctrl)
	mockRabbitMQ := mocks.NewMockIRabbitMQ(ctrl)
	mockMemcached := mocks.NewMockMencacher(ctrl)
	var messageChannel = make(chan amqp.Delivery, 1)
	var returnChannel = make(chan error, 1)

	userService := services.NewUserService(mockUserRepository, mockRabbitMQ, messageChannel, returnChannel, mockMemcached)
	return userService, mockUserRepository, mockRabbitMQ, mockMemcached
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
	CreatedAt: createdAt,
	UpdatedAt: nil,
}

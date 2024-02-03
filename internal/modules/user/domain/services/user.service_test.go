package services_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/mocks"
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

var password = "hgGFHJ654*"
var user, err = entities.NewUser("John", "Doe", "123.356.229-00", "john.doe@example.com", "hgGFHJ654*")

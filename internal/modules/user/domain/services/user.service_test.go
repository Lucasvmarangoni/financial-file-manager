package services_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/mocks"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
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

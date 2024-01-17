package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Create(t *testing.T) {
	userService, _, mockRabbitMQ := prepare(t)

	mockRabbitMQ.EXPECT().
		Publish(gomock.Any(), "application/json", gomock.Any(), gomock.Any()).
		Return().
		Times(1)

	err := userService.Create("John", "Doe", "123.356.229-00", "john.doe@example.com", "hjH**g54gHç")
	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}
	assert.Nil(t, err)
}

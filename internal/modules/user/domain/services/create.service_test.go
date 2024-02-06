package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Create(t *testing.T) {
	userService, _, mockRabbitMQ, _ := prepare(t)

	t.Run("Should returned nil and publish in rabbitMQ queue when valid user is provided", func(t *testing.T) {

		mockRabbitMQ.EXPECT().
			Publish(gomock.Any(), "application/json", gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		go func() {
			userService.ReturnChannel <- nil
		}()

		err := userService.Create("John", "Doe", "123.356.229-00", "john.doe@example.com", "hjH**g54gHç")
		if err != nil {
			t.Errorf("Create returned an error: %v", err)
		}
		assert.Nil(t, err)
	})

	t.Run("Should return an error when invalid param is provided", func(t *testing.T) {

		invalid_cpf := "12335622900"
		err := userService.Create("John", "Doe", invalid_cpf, "john.doe@example.com", "hjH**g54gHç")
		assert.Error(t, err)
	})
}

func BenchmarkUserService_Create(b *testing.B) {
	userService, _, mockRabbitMQ, _ := prepare(b)

	mockRabbitMQ.EXPECT().
		Publish(gomock.Any(), "application/json", gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	for i := 0; i < b.N; i++ {
		go func() {
			userService.ReturnChannel <- nil
		}()
		_ = userService.Create("John", "Doe", "123.356.229-00", "john.doe@example.com", "hjH**g54gHç")
	}

}

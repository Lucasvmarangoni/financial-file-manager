package services_test

import (
	"strings"
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Authn(t *testing.T) {
	userService, mockRepo, _, mockMemcached := prepare(t)

	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()

	t.Run("Should return token when valid cpf is provided", func(t *testing.T) {
		cpf := "235.411.314-00"

		user, _ := entities.NewUser("John", "Doe", "123.356.229-00", "john.doe@example.com", password)

		mockRepo.EXPECT().
			FindByCpf(entities.Hash(cpf), gomock.Any()).
			Return(user, nil).Times(1)

		token, err := userService.Authn(cpf, password, config.GetTokenAuth(), 3600)

		assert.NoError(t, err)
		assert.NotEmpty(t, token, "The token should not be empty")

		parts := strings.Split(token, ".")
		assert.Equal(t, 3, len(parts), "The token should have three parts separated by dots")
	})

	t.Run("Should return token when valid email is provided", func(t *testing.T) {
		emailToFind := "john.doe@example.com"

		mockRepo.EXPECT().
			FindByEmail(entities.Hash(emailToFind), gomock.Any()).
			Return(user, nil).Times(1)

		token, err := userService.Authn(emailToFind, password, config.GetTokenAuth(), 3600)

		assert.NoError(t, err)
		assert.NotEmpty(t, token, "The token should not be empty")

		parts := strings.Split(token, ".")
		assert.Equal(t, 3, len(parts), "The token should have three parts separated by dots")
	})
}

func BenchmarkUserService_Authn(b *testing.B) {
	userService, mockRepo, _, mockMemcached := prepare(b)

	user, _ := entities.NewUser("John", "Doe", "123.356.229-00", "john.doe@example.com", password)

	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()

	emailToFind := "john.doe@example.com"

	mockRepo.EXPECT().
		FindByEmail(entities.Hash(emailToFind), gomock.Any()).
		Return(user, nil).AnyTimes()

	for i := 0; i < b.N; i++ {
		_, _ = userService.Authn(emailToFind, password, config.GetTokenAuth(), 3600)
	}
}

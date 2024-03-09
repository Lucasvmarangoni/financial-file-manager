package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/security"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_FindByEmail(t *testing.T) {
	userService, mockRepo, _, mockMemcached := prepare(t)
	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()

	emailToFind := "john.doe@example.com"
	invalid_email := "invalid@example.com"

	t.Run("Should returned a valid user when valid and existing email is provided", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail(entities.Hash(emailToFind), gomock.Any()).
			Return(user, nil).Times(1)

		foundUser, err := userService.FindByEmail(entities.Hash(emailToFind), context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.Email, foundUser.Email)
	})

	t.Run("Should returned a error when invalid email is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindByEmail(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("User not found")).Times(1)

		_, err := userService.FindByEmail(invalid_email, context.Background())
		assert.Error(t, err)
	})
}

func TestUserService_FindByCpf(t *testing.T) {
	userService, mockRepo, _, mockMemcached := prepare(t)
	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()

	cpfToFind := "123.356.229-00"
	invalid_cpf := "123356229-00"

	t.Run("Should returned a valid user when valid and existing cpf is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindByCpf(entities.Hash(cpfToFind), gomock.Any()).
			Return(user, nil).Times(1)

		foundUser, err := userService.FindByCpf(entities.Hash(cpfToFind), context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.Email, foundUser.Email)
	})

	t.Run("Should returned a error when invalid cpf is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindByCpf(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("User not found")).Times(1)

		_, err := userService.FindByCpf(invalid_cpf, context.Background())

		assert.Error(t, err)
	})
}

func TestUserService_FindById(t *testing.T) {
	userService, mockRepo, _, mockMemcached := prepare(t)
	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()
	invalid_id := "52c599f83-4fd9-bfd6-e532918f7b13"

	t.Run("Should returned a valid user when valid and existing ID is provided", func(t *testing.T) {

		user, err := entities.NewUser("John", "Doe", "123.356.229-00", "john.doe@example.com", "hgGFHJ654*")
		aes_key := config.GetEnvString("security", "aes_key")

		encryptedEmail, _ := security.Encrypt(user.Email, aes_key)
		encryptedCPF, _ := security.Encrypt(user.CPF, aes_key)
		encryptedLastName, _ := security.Encrypt(user.LastName, aes_key)

		user.Email = encryptedEmail
		user.CPF = encryptedCPF
		user.LastName = encryptedLastName

		mockRepo.EXPECT().
			FindById(user.ID, gomock.Any()).
			Return(user, nil).Times(1)

		foundUser, err := userService.FindById(user.ID.String(), context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, foundUser)

		assert.Equal(t, user.ID, foundUser.ID)
	})

	t.Run("Should returned a error invalid id is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindById(invalid_id, gomock.Any()).
			Return(nil, errors.New("User not found")).Times(1)

		_, err := userService.FindById(invalid_id, context.Background())

		assert.Error(t, err)
	})
}

func BenchmarkUserService_FindBeEmail(b *testing.B) {
	userService, mockRepo, _, mockMemcached := prepare(b)

	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()

	emailToFind := "john.doe@example.com"
	mockRepo.EXPECT().
		FindByEmail(entities.Hash(emailToFind), gomock.Any()).
		Return(user, nil).AnyTimes()

	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes().Return(nil, errors.New("key not found"))

	for i := 0; i < b.N; i++ {
		_, _ = userService.FindByEmail(entities.Hash(emailToFind), context.Background())
	}
}

func BenchmarkUserService_FindBeCpf(b *testing.B) {
	userService, mockRepo, _, mockMemcached := prepare(b)

	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()

	cpfToFind := "123.356.229-00"
	mockRepo.EXPECT().
		FindByCpf(entities.Hash(cpfToFind), gomock.Any()).
		Return(user, nil).AnyTimes()

	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes().Return(nil, errors.New("key not found"))

	for i := 0; i < b.N; i++ {
		_, _ = userService.FindByCpf(entities.Hash(cpfToFind), context.Background())
	}
}

func BenchmarkUserService_FindBeId(b *testing.B) {
	userService, mockRepo, _, mockMemcached := prepare(b)

	mockRepo.EXPECT().
		FindById(user.ID, gomock.Any()).
		Return(user, nil).AnyTimes()

	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes().Return(nil, errors.New("key not found"))

	for i := 0; i < b.N; i++ {
		_, _ = userService.FindById(user.ID.String(), context.Background())
	}
}

package services_test

import (
	"context"
	"errors"
	"testing"

	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_FindByEmail(t *testing.T) {
	userService, mockRepo, _ := prepare(t)

	emailToFind := "john.doe@example.com"
	invalid_email := "invalid@example.com"
	t.Run("Should returned a valid user when valid and existing email is provided", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail(emailToFind, gomock.Any()).
			Return(user, nil).Times(1)

		foundUser, err := userService.FindByEmail(emailToFind, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.Email, foundUser.Email)
	})

	t.Run("Should returned a error when invalid email is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindByEmail(invalid_email, gomock.Any()).
			Return(nil, errors.New("User not found")).Times(1)

		_, err := userService.FindByEmail(invalid_email, context.Background())

		assert.Error(t, err)
	})
}

func TestUserService_FindByCpf(t *testing.T) {
	userService, mockRepo, _ := prepare(t)

	cpfToFind := "123.356.229-00"
	invalid_cpf := "123356229-00"

	t.Run("Should returned a valid user when valid and existing cpf is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindByCpf(cpfToFind, gomock.Any()).
			Return(user, nil).Times(1)

		foundUser, err := userService.FindByCpf(cpfToFind, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.Email, foundUser.Email)

	})

	t.Run("Should returned a error when invalid cpf is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindByCpf(invalid_cpf, gomock.Any()).
			Return(nil, errors.New("User not found")).Times(1)

		_, err := userService.FindByCpf(invalid_cpf, context.Background())

		assert.Error(t, err)
	})
}


func TestUserService_FindById(t *testing.T) {
	userService, mockRepo, _ := prepare(t)
	id, _ := pkg_entities.ParseID("52c599f3-af83-4fd9-bfd6-e532918f7b13")
	idToFind := id
	invalid_id := "52c599f83-4fd9-bfd6-e532918f7b13"

	t.Run("Should returned a valid user when valid and existing ID is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindById(idToFind, gomock.Any()).
			Return(user, nil).Times(1)

		foundUser, err := userService.FindById(idToFind.String(), context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.Email, foundUser.Email)
	})

	t.Run("Should returned a error invalid id is provided", func(t *testing.T) {

		mockRepo.EXPECT().
			FindById(invalid_id, gomock.Any()).
			Return(nil, errors.New("User not found")).Times(1)

		_, err := userService.FindById(invalid_id, context.Background())

		assert.Error(t, err)
	})
}

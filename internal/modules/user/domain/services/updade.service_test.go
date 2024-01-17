package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Update(t *testing.T) {
	userService, mockRepo, _ := prepare(t)
	id, _ := pkg_entities.ParseID("52c599f3-af83-4fd9-bfd6-e532918f7b13")
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

	t.Run("Should updated user when valid params is provided", func(t *testing.T) {
		mockRepo.EXPECT().
			FindById(id, gomock.Any()).
			Return(user, nil).Times(1)

		new_lastname := "NewLastName"
		new_email := "new-email@example.com"

		updated_user := &entities.User{
			ID:        id,
			Name:      "John",
			LastName:  new_lastname,
			CPF:       "123.356.229-00",
			Email:     new_email,
			Password:  user.Password,
			Admin:     false,
			CreatedAt: createdAt,
			UpdatedAt: []time.Time{time.Now()},
		}		

		mockRepo.EXPECT().
    Update(gomock.Any(), gomock.Any()). 
    Do(func(user *entities.User, _ context.Context) {     
        assert.Equal(t, updated_user.ID, user.ID)
        assert.Equal(t, updated_user.Name, user.Name)
        assert.Equal(t, updated_user.LastName, user.LastName)
        assert.Equal(t, updated_user.CPF, user.CPF)
        assert.Equal(t, updated_user.Email, user.Email)
        assert.Equal(t, updated_user.Admin, user.Admin)
        assert.Equal(t, updated_user.CreatedAt, user.CreatedAt)        
    }).
    Return(updated_user, nil).Times(1)

		err := userService.Update(id.String(), "", new_lastname, new_email, "")

		assert.Nil(t, err)
	})

}

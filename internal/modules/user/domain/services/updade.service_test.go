package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/security"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Update(t *testing.T) {

	user, err := entities.NewUser("John", "Doe", "123.356.229-00", "john.doe@example.com", "hgGFHJ654*")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	userService, mockRepo, _, mockMemcached := prepare(t)
	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()
	aes_key := config.ReadSecretString(config.GetEnvString("security", "aes_key"))

	t.Run("Should updated user when valid params is provided", func(t *testing.T) {

		encryptedEmail, _ := security.Encrypt(user.Email, aes_key)
		encryptedCPF, _ := security.Encrypt(user.CPF, aes_key)
		encryptedLastName, _ := security.Encrypt(user.LastName, aes_key)

		user.Email = encryptedEmail
		user.CPF = encryptedCPF
		user.LastName = encryptedLastName

		mockRepo.EXPECT().
			FindById(user.ID, gomock.Any()).
			Return(user, nil).Times(1)

		new_lastname := "NewLastName"
		new_email := "new-email@example.com"

		updated_user := &entities.User{
			ID:        user.ID,
			Name:      "John",
			LastName:  new_lastname,
			CPF:       "123.356.229-00",
			Email:     new_email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdateLog: []entities.UpdateLog{
				{
					Timestamp: time.Now(),
					OldValues: map[string]interface{}{
						"LastName": user.LastName,
						"Email":    user.Email,
					},
				},
			},
		}

		mockRepo.EXPECT().
			Update(gomock.Any(), gomock.Any()).
			Do(func(user *entities.User, _ context.Context) {
				assert.Equal(t, updated_user.ID, user.ID)
				assert.Equal(t, updated_user.Name, user.Name)

				assert.NotEmpty(t, user.LastName)
				assert.NotEmpty(t, user.CPF)
				assert.NotEmpty(t, user.Email)
				assert.Equal(t, updated_user.CreatedAt, user.CreatedAt)
			}).
			Return(nil).Times(1)

		err := userService.Update(user.ID.String(), "", new_lastname, new_email, password, "")
		assert.Nil(t, err)
	})

}

func BenchmarkUserService_Update(b *testing.B) {
	userService, mockRepo, _, mockMemcached := prepare(b)
	mockMemcached.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
	mockMemcached.EXPECT().Get(gomock.Any()).AnyTimes()

	mockRepo.EXPECT().
		FindById(user.ID, gomock.Any()).
		Return(user, nil).AnyTimes()

	new_lastname := "NewLastName"
	new_email := "new-email@example.com"

	updated_user := &entities.User{
		ID:        user.ID,
		Name:      "John",
		LastName:  new_lastname,
		CPF:       "123.356.229-00",
		Email:     new_email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdateLog: []entities.UpdateLog{
			{
				Timestamp: time.Now(),
				OldValues: map[string]interface{}{
					"LastName": user.LastName,
					"Email":    user.Email,
				},
			},
		},
	}

	mockRepo.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Do(func(user *entities.User, _ context.Context) {
			assert.Equal(b, updated_user.ID, user.ID)
			assert.Equal(b, updated_user.Name, user.Name)
			assert.Equal(b, updated_user.LastName, user.LastName)
			assert.Equal(b, updated_user.CPF, user.CPF)
			assert.Equal(b, updated_user.Email, user.Email)
			assert.Equal(b, updated_user.CreatedAt, user.CreatedAt)
		}).
		Return(nil).AnyTimes()

	for i := 0; i < b.N; i++ {
		_ = userService.Update(user.ID.String(), "", new_lastname, new_email, user.Password, "")
	}
}

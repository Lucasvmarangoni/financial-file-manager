package services_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_AuthzAdmin(t *testing.T) {
	userService, mockRepo, _ := prepare(t)
	adminId, _ := pkg_entities.ParseID("52c599f3-af83-4fd9-bfd6-e532918f7b23")

	admin := &entities.User{
		ID:        adminId,
		Name:      "John",
		LastName:  "Doe",
		CPF:       "123.356.239-00",
		Email:     "john.do2e@example.com",
		Password:  password,
		Admin:     true,
		CreatedAt: createdAt,
		UpdatedAt: nil,
	}

	t.Run("AuthzAdmin", func(t *testing.T) {
		mockRepo.EXPECT().
			FindById(gomock.Eq(adminId), gomock.Any()).
			Return(admin, nil).Times(1)

		mockRepo.EXPECT().
			FindById(gomock.Eq(id), gomock.Any()).
			Return(user, nil).Times(1)

		mockRepo.EXPECT().
			ToggleAdmin(gomock.Eq(id.String()), gomock.Any()).
			Return(nil).Times(1)

		err := userService.AuthzAdmin(id.String())

		assert.Nil(t, err)
	})
}

func BenchmarkUserService_AuthzAdmin(b *testing.B) {
	userService, mockRepo, _ := prepare(b)
	adminId, _ := pkg_entities.ParseID("52c599f3-af83-4fd9-bfd6-e532918f7b23")

	admin := &entities.User{
		ID:        adminId,
		Name:      "John",
		LastName:  "Doe",
		CPF:       "123.356.239-00",
		Email:     "john.do2e@example.com",
		Password:  password,
		Admin:     true,
		CreatedAt: createdAt,
		UpdatedAt: nil,
	}

	mockRepo.EXPECT().
		FindById(gomock.Eq(adminId), gomock.Any()).
		Return(admin, nil).AnyTimes()

	mockRepo.EXPECT().
		FindById(gomock.Eq(id), gomock.Any()).
		Return(user, nil).AnyTimes()

	mockRepo.EXPECT().
		ToggleAdmin(gomock.Eq(id.String()), gomock.Any()).
		Return(nil).AnyTimes()

	for i := 0; i < b.N; i++ {
		_ = userService.AuthzAdmin(id.String())
	}
}

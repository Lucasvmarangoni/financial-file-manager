package services_test

import (
	"testing"

	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Delete(t *testing.T) {
	userService, mockRepo, _ := prepare(t)
	id, _ := pkg_entities.ParseID("52c599f3-af83-4fd9-bfd6-e532918f7b13")
	t.Run("Should delete user", func(t *testing.T) {
		mockRepo.EXPECT().
			Delete(gomock.Eq(id.String()), gomock.Any()).
			Return(nil).Times(1)

		err := userService.Delete(id.String())

		assert.Nil(t, err)
	})
}

package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Delete(t *testing.T) {
	userService, mockRepo, _ := prepare(t)
	t.Run("Should delete user", func(t *testing.T) {
		mockRepo.EXPECT().
			Delete(gomock.Eq(id.String()), gomock.Any()).
			Return(nil).Times(1)

		err := userService.Delete(id.String())

		assert.Nil(t, err)
	})
}

func BenchmarkUserService_Delete(b *testing.B) {
	userService, mockRepo, _ := prepare(b)

	mockRepo.EXPECT().
		Delete(gomock.Eq(id.String()), gomock.Any()).
		Return(nil).AnyTimes()

	for i := 0; i < b.N; i++ {
		_ = userService.Delete(id.String())
	}
}

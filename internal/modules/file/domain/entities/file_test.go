package entities_test

import (
	"log"
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	entities_pkg "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/google/uuid"
)

func init() {
	err := godotenv.Load("../../../../..//config/.env.default")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestNewFile(t *testing.T) {

	typ := "contract"
	user := entities_pkg.NewID()
	now := time.Now()
	t.Run("should return a new file when provided valid params", func(t *testing.T) {

		file, err := entities.NewFile(typ, user, nil, false)

		require.NotNil(t, file)
		require.Nil(t, err)
		assert.Equal(t, typ, file.Type)
		assert.True(t, file.CreatedAt.After(now) || file.CreatedAt.Equal(now))
		assert.Equal(t, user, file.User)
		assert.NotEmpty(t, file.ID)
	})

	t.Run("should return error when invalid param type is provided", func(t *testing.T) {

		typ = "-"

		file, err := entities.NewFile(typ, user, nil, false)

		require.NotNil(t, err)
		require.Nil(t, file)
	})

	t.Run("should return error when invalid param user is provided", func(t *testing.T) {

		user = uuid.Nil

		file, err := entities.NewFile(typ, user, nil, false)

		require.NotNil(t, err)
		require.Nil(t, file)
	})
}

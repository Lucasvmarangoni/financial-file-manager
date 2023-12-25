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
	customer := entities_pkg.NewID()
	now := time.Now()
	t.Run("should return a new file when provided valid params", func(t *testing.T) {

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, file)
		require.Nil(t, err)
		assert.Equal(t, typ, file.Type)
		assert.True(t, file.CreatedAt.After(now) || file.CreatedAt.Equal(now))
		assert.Equal(t, customer, file.Customer)
		assert.NotEmpty(t, file.ID)
	})

	t.Run("should return error when invalid param type is provided", func(t *testing.T) {

		typ = "-"

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, err)
		require.Nil(t, file)
	})

	t.Run("should return error when invalid param Customer is provided", func(t *testing.T) {

		customer = uuid.Nil

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, err)
		require.Nil(t, file)
	})
}

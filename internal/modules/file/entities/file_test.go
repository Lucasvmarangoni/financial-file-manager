package entities_test

import (
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	t.Run("should return error when type is invalid", func(t *testing.T) {
		file := entities.File{
			ID:        "test-id",
			Type:      "invalid",
			CreatedAt: time.Now(),
			Customer:  "test-customer",
		}

		err := file.Validate()
		assert.Error(t, err)
		assert.Equal(t, "invalid type", err.Error())
	})

	t.Run("should return error when struct is invalid", func(t *testing.T) {
		file := entities.File{
			ID:        "test-id",
			Type:      "contract",
			CreatedAt: time.Now(),
		}

		err := file.Validate()
		assert.Error(t, err)
	})

	t.Run("should return nil when file is valid", func(t *testing.T) {
		file := entities.File{
			ID:        "123e4567-e89b-12d3-a456-426614174000",
			Type:      "contract",
			CreatedAt: time.Now(),
			Customer:  "test-customer",
		}

		err := file.Validate()
		assert.NoError(t, err)
	})
}

func TestNewFile(t *testing.T) {
	t.Run("should return a new file with the provided parameters", func(t *testing.T) {
		typ := "contract"
		createdAt := time.Now()
		customer := "test-customer"

		file, err := entities.NewFile(typ, createdAt, customer)

		require.NotNil(t, file)
		require.Nil(t, err)
		assert.Equal(t, typ, file.Type)
		assert.Equal(t, createdAt, file.CreatedAt)
		assert.Equal(t, customer, file.Customer)
		assert.NotEmpty(t, file.ID)
	})
}

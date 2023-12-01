package entities_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	consts "github.com/Lucasvmarangoni/financial-file-manager/internal/common/const"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../../common/config/.env.default")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
		assert.Equal(t, fmt.Sprintf("invalid type, must be: %v", consts.FileTypes()), err.Error())
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
		
		customer := "test-customer"
		now := time.Now()

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, file)
		require.Nil(t, err)
		assert.Equal(t, typ, file.Type)
		assert.True(t, file.CreatedAt.After(now) || file.CreatedAt.Equal(now))
		assert.Equal(t, customer, file.Customer)
		assert.NotEmpty(t, file.ID)
	})
}

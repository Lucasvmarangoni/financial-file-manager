package entities_test

import (
	"log"
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	entities_pkg "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../../../..//config/.env.default")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func TestNewExtract(t *testing.T) {
	typ := "contract"
	customer := entities_pkg.NewID()

	file, err := entities.NewFile(typ, customer)

	require.NotNil(t, file)
	require.Nil(t, err)

	account := 1
	value := 20.20
	category := "deposit"
	method := "debit"
	location := "test-location"
	t.Run("should return a new extract when valid params are provided", func(t *testing.T) {

		extract, err := entities.NewExtract(
			*file,
			account,
			value,
			category,
			method,
			location,
			uuid.Nil,
		)

		require.NotNil(t, extract)
		require.Nil(t, err)
		assert.Equal(t, file, extract.File)
		assert.Equal(t, account, extract.Account)
		assert.Equal(t, value, extract.Value)
		assert.Equal(t, category, extract.Category)
		assert.Equal(t, method, extract.Method)
		assert.Equal(t, location, extract.Location)
	})

	t.Run("should return error when invalid param account is provided", func(t *testing.T) {

		account = 0

		extract, err := entities.NewExtract(
			*file,
			account,
			value,
			category,
			method,
			location,
			uuid.Nil,
		)

		require.NotNil(t, err)
		require.Nil(t, extract)
	})

	t.Run("should return error when invalid param value is provided", func(t *testing.T) {

		value = 0

		extract, err := entities.NewExtract(
			*file,
			account,
			value,
			category,
			method,
			location,
			uuid.Nil,
		)

		require.NotNil(t, err)
		require.Nil(t, extract)
	})

	t.Run("should return error when invalid param category is provided", func(t *testing.T) {

		category = "-"

		extract, err := entities.NewExtract(
			*file,
			account,
			value,
			category,
			method,
			location,
			uuid.Nil,
		)

		require.NotNil(t, err)
		require.Nil(t, extract)
	})

	t.Run("should return error when invalid param method is provided", func(t *testing.T) {

		method = "-"

		extract, err := entities.NewExtract(
			*file,
			account,
			value,
			category,
			method,
			location,
			uuid.Nil,
		)

		require.NotNil(t, err)
		require.Nil(t, extract)
	})

	t.Run("should return error when invalid param location is provided", func(t *testing.T) {

		location = ""

		extract, err := entities.NewExtract(
			*file,
			account,
			value,
			category,
			method,
			location,
			uuid.Nil,
		)

		require.NotNil(t, err)
		require.Nil(t, extract)
	})
}

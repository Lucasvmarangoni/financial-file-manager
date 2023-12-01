package entities_test

import (
	"fmt"
	"log"
	"testing"

	consts "github.com/Lucasvmarangoni/financial-file-manager/internal/common/const"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../../../common/config/.env.default")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestExtractValidate(t *testing.T) {

	t.Run("should return error when account is lass than or equal 0", func(t *testing.T) {
		typ := "contract"
		customer := "test-customer"

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, file)
		require.Nil(t, err)

		extract := entities.Extract{
			File:     *file,
			Account:  0,
			Value:    20.0,
			Category: "deposit",
			Method:   "debit",
			Location: "test-location",
		}

		err = extract.Validate()
		assert.Error(t, err)
		assert.Equal(t, "Account needs to be greater than 0", err.Error())
	})

	t.Run("should return error when invalid method is provided", func(t *testing.T) {
		typ := "contract"
		customer := "test-customer"

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, file)
		require.Nil(t, err)

		extract := entities.Extract{
			File:     *file,
			Account:  1,
			Value:    20.0,
			Category: "deposit",
			Method:   "invalid",
			Location: "test-location",
		}

		err = extract.Validate()
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf("Need a valid method: %v", consts.Method()), err.Error())
	})
}

func TextNewExtract(t *testing.T) {

	t.Run("should return a new extract", func(t *testing.T) {

		typ := "contract"
		customer := "test-customer"

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, file)
		require.Nil(t, err)

		account := 0
		value := 20.20
		category := "deposit"
		method := "debit"
		location := "test-location"

		extract, err := entities.NewExtract(
			*file,
			account,
			value,
			category,
			method,
			location,
			nil,
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
}

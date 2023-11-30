package entities_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/common/const"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractValidate(t *testing.T) {

	t.Run("should return error when account is lass than or equal 0", func(t *testing.T) {
		extract := entities.Extract{
			File: entities.File{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				Type:      "extract",
				CreatedAt: time.Now(),
				Customer:  "test-customer",
			},
			Account:  0,
			Value:    20.0,
			Category: "deposit",
			Method:   "debit",
			Location: "test-location",
		}

		err := extract.Validate()
		assert.Error(t, err)
		assert.Equal(t, "Account needs to be greater than 0", err.Error())
	})

	t.Run("should return error when invalid method is provided", func(t *testing.T) {
		extract := entities.Extract{
			File: entities.File{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				Type:      "extract",
				CreatedAt: time.Now(),
				Customer:  "test-customer",
			},
			Account:  1,
			Value:    20.0,
			Category: "deposit",
			Method:   "invalid",
			Location: "test-location",
		}

		err := extract.Validate()
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf("Need a valid method: %v", consts.Method()), err.Error())
	})
}

func TextNewExtract(t *testing.T) {

	t.Run("should return a new extract", func(t *testing.T) {

		file := entities.File{
			ID:        "123e4567-e89b-12d3-a456-426614174000",
			Type:      "extract",
			CreatedAt: time.Now(),
			Customer:  "test-customer",
		}

		account := 0
		value := 20.20
		category := "deposit"
		method := "debit"
		location := "test-location"

		extract, err := entities.NewExtract(
			file,
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

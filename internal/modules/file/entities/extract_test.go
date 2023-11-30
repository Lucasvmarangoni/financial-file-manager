package entities_test

import (
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
	"github.com/stretchr/testify/assert"
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
			Method:   "card",
			Location: "test-location",
		}

		err := extract.Validate()
		assert.Error(t, err)
		assert.Equal(t, "Account needs to be greater than 0", err.Error())
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

		account:=  0
		value:=    20.20
		category:= "deposit"
		method:=   "card"
		location:= "test-location"

		extract := entities.NewExtract(
			file, 
			account, 
			value, 
			category, 
			method, 
			location, 
			nil,
		)
		
		assert.Equal(t, file, extract.File)
		assert.Equal(t, account, extract.Account)
		assert.Equal(t, value, extract.Value)
		assert.Equal(t, category, extract.Category)
		assert.Equal(t, method, extract.Method)
		assert.Equal(t, location, extract.Location)
	})
}

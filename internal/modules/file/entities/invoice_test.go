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

func TestInvoiceValidate(t *testing.T) {

	t.Run("should return error when invalid method is provided", func(t *testing.T) {
		dueDate, _ := time.Parse(time.RFC3339, "2022-03-14T09:26:22.123456789-07:00")
		methods := consts.Method()

		invoice := entities.Invoice{
			File: entities.File{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				Type:      "invoice",
				CreatedAt: time.Now(),
				Customer:  "test-customer",
			},
			DueDate: dueDate,
			Value:   12.0,
			Method:  "invalid",
		}

		err := invoice.Validate()
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf("Need a valid method: %v", methods), err.Error())
	})
}

func TextNewInvoice(t *testing.T) {

	t.Run("should return a new invoice", func(t *testing.T) {

		file := entities.File{
			ID:        "123e4567-e89b-12d3-a456-426614174000",
			Type:      "invoice",
			CreatedAt: time.Now(),
			Customer:  "test-customer",
		}
		dueDate, _ := time.Parse(time.RFC3339, "2022-03-14T09:26:22.123456789-07:00")
		value := 29.0
		method := "debit"

		invoice, err := entities.NewInvoice(file, dueDate, value, method, nil)

		require.NotNil(t, invoice)
		require.Nil(t, err)
		assert.Equal(t, file, invoice.File)
		assert.Equal(t, dueDate, invoice.DueDate)
		assert.Equal(t, value, invoice.Value)
		assert.Equal(t, method, invoice.Method)		
	})
}

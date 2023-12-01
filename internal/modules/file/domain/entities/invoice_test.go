package entities_test

import (
	"fmt"
	"testing"
	"time"

	consts "github.com/Lucasvmarangoni/financial-file-manager/internal/common/const"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvoiceValidate(t *testing.T) {

	t.Run("should return error when invalid method is provided", func(t *testing.T) {
		dueDate, _ := time.Parse(time.RFC3339, "2022-03-14T09:26:22.123456789-07:00")
		methods := consts.Method()
		typ := "contract"
		customer := "test-customer"

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, file)
		require.Nil(t, err)

		invoice := entities.Invoice{
			File:    *file,
			DueDate: dueDate,
			Value:   12.0,
			Method:  "invalid",
		}

		err = invoice.Validate()
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf("Need a valid method: %v", methods), err.Error())
	})
}

func TextNewInvoice(t *testing.T) {

	t.Run("should return a new invoice", func(t *testing.T) {

		typ := "contract"
		customer := "test-customer"

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, file)
		require.Nil(t, err)
		dueDate, _ := time.Parse(time.RFC3339, "2022-03-14T09:26:22.123456789-07:00")
		value := 29.0
		method := "debit"

		invoice, err := entities.NewInvoice(*file, dueDate, value, method, nil)

		require.NotNil(t, invoice)
		require.Nil(t, err)
		assert.Equal(t, file, invoice.File)
		assert.Equal(t, dueDate, invoice.DueDate)
		assert.Equal(t, value, invoice.Value)
		assert.Equal(t, method, invoice.Method)
	})
}

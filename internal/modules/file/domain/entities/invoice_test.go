package entities_test

import (
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	entities_pkg "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInvoice(t *testing.T) {

	typ := "contract"
	customer := entities_pkg.NewID()

	file, err := entities.NewFile(typ, customer, nil, false)

	require.NotNil(t, file)
	require.Nil(t, err)
	dueDate, _ := time.Parse(time.RFC3339, "2022-03-14T09:26:22.123456789-07:00")
	value := 29.0
	method := "debit"

	t.Run("should return a new invoice when valid params are provided", func(t *testing.T) {

		invoice, err := entities.NewInvoice(*file, dueDate, value, method, uuid.Nil)

		require.NotNil(t, invoice)
		require.Nil(t, err)
		assert.Equal(t, file, invoice.File)
		assert.Equal(t, dueDate, invoice.DueDate)
		assert.Equal(t, value, invoice.Value)
		assert.Equal(t, method, invoice.Method)
	})

	t.Run("should return error when invalid params dueDate is provided", func(t *testing.T) {

		var dueDate time.Time

		invoice, err := entities.NewInvoice(*file, dueDate, value, method, uuid.Nil)

		require.NotNil(t, err)
		require.Nil(t, invoice)
	})

	t.Run("should return error when invalid params value is provided", func(t *testing.T) {

		value = 0

		invoice, err := entities.NewInvoice(*file, dueDate, value, method, uuid.Nil)

		require.NotNil(t, err)
		require.Nil(t, invoice)
	})

	t.Run("should return error when invalid params method is provided", func(t *testing.T) {

		method = "-"

		invoice, err := entities.NewInvoice(*file, dueDate, value, method, uuid.Nil)

		require.NotNil(t, err)
		require.Nil(t, invoice)
	})
}

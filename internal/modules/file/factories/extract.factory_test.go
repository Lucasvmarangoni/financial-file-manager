package factories_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/factories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TextExtractFactory(t *testing.T) {

	typ := "invoice"
	customer := "test-customer"
	account := 0
	value := 20.0
	category := "deposit"
	method := "debit"
	location := "test-location"

	t.Run("should return a new extract", func(t *testing.T) {

		extract, err := factories.ExtractFactory(
			typ,
			customer,
			account,
			value,
			category,
			method,
			location,
			nil,
		)
		require.NotNil(t, extract)
		require.Nil(t, err)
	})

	t.Run("should return error when invalid type is provided", func(t *testing.T) {

		typ = "invalid"

		extract, err := factories.ExtractFactory(
			typ,
			customer,
			account,
			value,
			category,
			method,
			location,
			nil,
		)
		require.Nil(t, extract)
		require.NotNil(t, err)
		assert.Equal(t, "invalid type, must be: extract, extract or invoice", err.Error())
	})
}

package factories_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/factories"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TextExtractFactory(t *testing.T) {

	typ := "invoice"
	user := pkg_entities.NewID()
	authorized := []pkg_entities.ID{pkg_entities.NewID(), pkg_entities.NewID()}
	account := 0
	value := 20.0
	category := "deposit"
	method := "debit"
	location := "test-location"

	t.Run("should return a new extract", func(t *testing.T) {

		extract, err := factories.ExtractFactory(
			typ,
			user,
			authorized,
			nil,
			account,
			value,
			category,
			method,
			location,
			uuid.Nil,
			false,
		)
		require.NotNil(t, extract)
		require.Nil(t, err)
	})

	t.Run("should return error when invalid type is provided", func(t *testing.T) {

		typ = "invalid"

		extract, err := factories.ExtractFactory(
			typ,
			user,
			authorized,
			nil,
			account,
			value,
			category,
			method,
			location,
			uuid.Nil,
			false,
		)
		require.Nil(t, extract)
		require.NotNil(t, err)
		assert.Equal(t, "invalid type, must be: extract, extract or invoice", err.Error())
	})
}

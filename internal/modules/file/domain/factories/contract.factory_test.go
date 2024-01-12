package factories_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/factories"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TextContractFactory(t *testing.T) {

	typ := "invoice"

	user := pkg_entities.NewID()
	title := "Test Title"
	parties := []string{"Party 1", "Party 2"}
	object := "Test Object"
	authorized := []pkg_entities.ID{pkg_entities.NewID(), pkg_entities.NewID()}

	t.Run("should return a new contract", func(t *testing.T) {

		contract, err := factories.ContractFactory(
			typ,
			user,
			authorized,
			nil,
			title,
			parties,
			object,
			nil,
			nil,
			false,
		)
		require.NotNil(t, contract)
		require.Nil(t, err)
	})

	t.Run("should return error when invalid type is provided", func(t *testing.T) {

		typ = "invalid"

		contract, err := factories.ContractFactory(
			typ,
			user,
			authorized,
			nil,
			title,
			parties,
			object,
			nil,
			nil,
			false,
		)
		require.Nil(t, contract)
		require.NotNil(t, err)
		assert.Equal(t, "invalid type, must be: contract, extract or invoice", err.Error())
	})
}

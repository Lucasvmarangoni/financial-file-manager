package entities_test

import (
	"testing"
	
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContractValidate(t *testing.T) {
	t.Run("should return error when title is empty", func(t *testing.T) {
		typ := "contract"
		customer := "test-customer"

		file, err := entities.NewFile(typ, customer)

		require.NotNil(t, file)
		require.Nil(t, err)

		contract := entities.Contract{
			File: *file,
			Title:   "",
			Parties: []string{"Party 1", "Party 2"},
			Object:  "Test Object",
		}

		err = contract.Validate()
		assert.Error(t, err)
		assert.Equal(t, "Need a title", err.Error())
	})	
}

func TestNewContract(t *testing.T) {
	typ := "contract"
	customer := "test-customer"

	file, err := entities.NewFile(typ, customer)

	require.NotNil(t, file)
	require.Nil(t, err)

	title := "Test Title"
	parties := []string{"Party 1", "Party 2"}
	object := "Test Object"
	extract := []string{"Extract 1", "Extract 2"}
	invoice := []string{"Invoice 1", "Invoice 2"}

	contract, err := entities.NewContract(*file, title, parties, object, extract, invoice)

	require.NotNil(t, contract)
	require.Nil(t, err)
	assert.Equal(t, *file, contract.File)
	assert.Equal(t, title, contract.Title)
	assert.Equal(t, parties, contract.Parties)
	assert.Equal(t, object, contract.Object)
	assert.Equal(t, extract, contract.Extract)
	assert.Equal(t, invoice, contract.Invoice)
}

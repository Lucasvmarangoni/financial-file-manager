package entities_test

import (
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContractValidate(t *testing.T) {
	t.Run("should return error when title is empty", func(t *testing.T) {
		contract := entities.Contract{
			File: entities.File{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				Type:      "contract",
				CreatedAt: time.Now(),
				Customer:  "test-customer",
			},
			Title:   "",
			Parties: []string{"Party 1", "Party 2"},
			Object:  "Test Object",
		}

		err := contract.Validate()
		assert.Error(t, err)
		assert.Equal(t, "Need a title", err.Error())
	})

	// Add more tests for other validation cases...
}

func TestNewContract(t *testing.T) {
	file := entities.File{
		ID:        "123e4567-e89b-12d3-a456-426614174000",
		Type:      "contract",
		CreatedAt: time.Now(),
		Customer:  "test-customer",
	}
	title := "Test Title"
	parties := []string{"Party 1", "Party 2"}
	object := "Test Object"
	extract := []string{"Extract 1", "Extract 2"}
	invoice := []string{"Invoice 1", "Invoice 2"}

	contract, err := entities.NewContract(file, title, parties, object, extract, invoice)

	require.NotNil(t, contract)
	require.Nil(t, err)
	assert.Equal(t, file, contract.File)
	assert.Equal(t, title, contract.Title)
	assert.Equal(t, parties, contract.Parties)
	assert.Equal(t, object, contract.Object)
	assert.Equal(t, extract, contract.Extract)
	assert.Equal(t, invoice, contract.Invoice)
}

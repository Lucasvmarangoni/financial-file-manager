package entities_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	entities_pkg "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewContract(t *testing.T) {

	typ := "contract"
	customer := entities_pkg.NewID()
	file, err := entities.NewFile(typ, customer)

	require.NotNil(t, file)
	require.Nil(t, err)

	title := "Test Title"
	parties := []string{"Party 1", "Party 2"}
	object := "Test Object"


	t.Run("Should return a new contract when valid params are provided", func(t *testing.T) {

		contract, err := entities.NewContract(file, title, parties, object, nil, nil)

		require.NotNil(t, contract)
		require.Nil(t, err)
		assert.Equal(t, *file, contract.File)
		assert.Equal(t, title, contract.Title)
		assert.Equal(t, parties, contract.Parties)
		assert.Equal(t, object, contract.Object)
	})

	t.Run("Should return error when invalid param title is provided", func(t *testing.T) {		
		title = "-"
	
		contract, err := entities.NewContract(file, title, parties, object, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, contract)		
	})

	t.Run("Should return error when invalid param parties is provided", func(t *testing.T) {		
		parties = []string{"Party 1"}
	
		contract, err := entities.NewContract(file, title, parties, object, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, contract)		
	})
}

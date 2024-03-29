package factories_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/factories"
	consts "github.com/Lucasvmarangoni/financial-file-manager/pkg/const"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TextInvoiceFactory(t *testing.T) {

	typ := "invoice"
	user := pkg_entities.NewID()
	authorized := []pkg_entities.ID{pkg_entities.NewID(), pkg_entities.NewID()}
	dueDate, _ := time.Parse(time.RFC3339, "2022-03-14T09:26:22.123456789-07:00")
	value := 12.0
	method := "invalid"

	t.Run("should return a new invoice", func(t *testing.T) {

		invoice, err := factories.InvoiceFactory(
			typ,
			user,
			authorized,
			nil,
			dueDate,
			value,
			method,
			uuid.Nil,
			false,
		)
		require.NotNil(t, invoice)
		require.Nil(t, err)
	})

	t.Run("should return error when invalid method is provided", func(t *testing.T) {

		method = "invalid"

		invoice, err := factories.InvoiceFactory(
			typ,
			user,
			authorized,
			nil,
			dueDate,
			value,
			method,
			uuid.Nil,
			false,
		)
		require.Nil(t, invoice)
		require.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("Need a valid method: %v", consts.Method()), err.Error())
	})
}

package factories

import (
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
)

func InvoiceFactory(
	typ string,
	customer pkg_entities.ID,
	dueDate time.Time,
	value float64,
	method string,
	contract pkg_entities.ID,
) (*entities.Invoice, error) {

	file, err := entities.NewFile(typ, customer)
	if err != nil {
		return nil, err
	}
	invoice, err := entities.NewInvoice(*file, dueDate, value, method, contract)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

package factories

import (
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
)

func InvoiceFactory(
	typ string,
	createdAt time.Time,
	customer string,
	dueDate time.Time,
	value float64,
	method string,
	contract *string,
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

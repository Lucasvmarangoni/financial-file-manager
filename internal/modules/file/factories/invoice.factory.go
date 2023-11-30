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
) *entities.Invoice {

	file := entities.NewFile(typ, createdAt, customer)
	invoice := entities.NewInvoice(*file, dueDate, value, method, contract)

	return invoice
}
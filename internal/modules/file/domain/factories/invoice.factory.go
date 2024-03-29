package factories

import (
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
)

func InvoiceFactory(
	typ string,
	user pkg_entities.ID,
	authorized []pkg_entities.ID,
	versions []pkg_entities.ID,
	dueDate time.Time,
	value float64,
	method string,
	contract pkg_entities.ID,
	archived bool,
) (*entities.Invoice, error) {

	file, err := entities.NewFile(typ, user, authorized, versions, archived)
	if err != nil {
		return nil, err
	}
	invoice, err := entities.NewInvoice(*file, dueDate, value, method, contract)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

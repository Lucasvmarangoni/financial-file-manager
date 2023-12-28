package factories

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
)

func ContractFactory(
	typ string,
	customer pkg_entities.ID,
	versions []pkg_entities.ID,
	title string,
	parties []string,
	object string,
	extract []pkg_entities.ID,
	invoice []pkg_entities.ID,
) (*entities.Contract, error) {

	file, err := entities.NewFile(typ, customer, versions)
	if err != nil {
		return nil, err
	}
	contract, err := entities.NewContract(file, title, parties, object, extract, invoice)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

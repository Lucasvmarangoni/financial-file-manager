package factories

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
)

func ContractFactory(
	typ string,
	user pkg_entities.ID,
	authorized []pkg_entities.ID,
	versions []pkg_entities.ID,
	title string,
	parties []string,
	object string,
	extract []pkg_entities.ID,
	invoice []pkg_entities.ID,
	archived bool,
) (*entities.Contract, error) {

	file, err := entities.NewFile(typ, user, authorized, versions, archived)
	if err != nil {
		return nil, err
	}
	contract, err := entities.NewContract(file, title, parties, object, extract, invoice)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

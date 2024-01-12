package factories

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
)

func ExtractFactory(
	typ string,
	user pkg_entities.ID,
	authorized []pkg_entities.ID,
	versions []pkg_entities.ID,
	account int,
	value float64,
	category string,
	method string,
	location string,
	contract pkg_entities.ID,
	archived bool,
) (*entities.Extract, error) {

	file, err := entities.NewFile(typ, user, authorized, versions, archived)
	if err != nil {
		return nil, err
	}
	extract, err := entities.NewExtract(*file, account, value, category, method, location, contract)
	if err != nil {
		return nil, err
	}
	return extract, nil
}

package factories

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
)

func ExtractFactory(
	typ string,
	customer pkg_entities.ID,
	account int,
	value float64,
	category string,
	method string,
	location string,
	contract pkg_entities.ID,
) (*entities.Extract, error) {

	file, err := entities.NewFile(typ, customer)
	if err != nil {
		return nil, err
	}
	extract, err := entities.NewExtract(*file, account, value, category, method, location, contract)
	if err != nil {
		return nil, err
	}
	return extract, nil
}

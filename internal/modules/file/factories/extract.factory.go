package factories

import (
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
)

func ExtractFactory(
	typ string,
	createdAt time.Time,
	customer string,
	account int,
	value float64,
	category string,
	method string,
	location string,
	contract *string,
) (*entities.Extract, error) {

	file, err := entities.NewFile(typ, createdAt, customer)
	if err != nil {
		return nil, err
	}
	extract, err := entities.NewExtract(*file, account, value, category, method, location, contract)
	if err != nil {
		return nil, err
	}
	return extract, nil
}

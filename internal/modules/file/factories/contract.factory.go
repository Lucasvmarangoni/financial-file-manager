package factories

import (
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/entities"
)

func ContractFactory(
	typ string, 
	createdAt time.Time, 
	customer string, 
	title string, 
	parties []string, 
	object string, 
	extract []string, 
	invoice []string,
	) (*entities.Contract, error) {

	file, err := entities.NewFile(typ, createdAt, customer)
	if err != nil {
		return nil, err
	}
	contract, err := entities.NewContract(*file, title, parties, object, extract, invoice)
	if err != nil {
		return nil, err
	}	
	return contract, nil
}

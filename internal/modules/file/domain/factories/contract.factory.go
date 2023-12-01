package factories

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
)

func ContractFactory(
	typ string, 	
	customer string, 
	title string, 
	parties []string, 
	object string, 
	extract []string, 
	invoice []string,
	) (*entities.Contract, error) {

	file, err := entities.NewFile(typ, customer)
	if err != nil {
		return nil, err
	}
	contract, err := entities.NewContract(*file, title, parties, object, extract, invoice)
	if err != nil {
		return nil, err
	}	
	return contract, nil
}

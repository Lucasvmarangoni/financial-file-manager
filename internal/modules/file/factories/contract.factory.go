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
	) *entities.Contract {

	file := entities.NewFile(typ, createdAt, customer)
	contract := entities.NewContract(*file, title, parties, object, extract, invoice)
	
	return contract
}

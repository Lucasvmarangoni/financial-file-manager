package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/const"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	pkg_errors "github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/lib"
	"github.com/asaskevich/govalidator"
)

/*

Para o requisito de manter versões do documento, analisar a ideia de
verificar se o ID foi passado na request, antes de atribuir.
Assim antes de enviar um documento que já existe, deve-se fazer uma
consulta por esse documento e enviar o novo com ID, para que o sistema identifique e
armazene de alguma forma que o histórico seja possível.

*/

type File struct {
	ID        entities.ID   `json:"document_id" valid:"notnull,required"`
	Type      string        `json:"type" valid:"notnull,required"`
	CreatedAt time.Time     `json:"created_at" valid:"-"`
	Customer  entities.ID   `json:"customer" valid:"-"`
	Versions  []entities.ID `json:"versions" valid:"-"`
}

func (f *File) Validate() error {

	fileTypes := consts.FileTypes()

	if _, err := entities.ParseID(f.ID.String()); err != nil {
		pkg_errors.IsInvalidError("ID", "Must be google uuid")
	}

	if _, err := entities.ParseID(f.Customer.String()); err != nil {
		pkg_errors.IsInvalidError("Customer", "Must be google uuid")
	}

	if !lib.MapVerifyString(fileTypes[:], strings.ToLower(f.Type)) {
		return pkg_errors.IsInvalidError("Type", fmt.Sprintf("Must be: %v", fileTypes))
	}

	_, err := govalidator.ValidateStruct(f)
	if err != nil {
		return err
	}
	return nil
}

func NewFile(typ string, customer entities.ID, versions []entities.ID) (*File, error) {

	file := File{
		Type:     typ,
		Customer: customer,
	}
	file.prepare()

	err := file.Validate()
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (f *File) prepare() {
	f.ID = entities.NewID()
	f.CreatedAt = time.Now()
}

package entities

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/common/const"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/common/lib"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/asaskevich/govalidator"
)

type File struct {
	ID        entities.ID `json:"document_id" valid:"-"`
	Type      string      `json:"type" valid:"notnull"`
	CreatedAt time.Time   `json:"created_at" valid:"-"`
	Customer  entities.ID      `json:"customer" valid:"notnull"`
}

func (f *File) Validate() error {

	fileTypes := consts.FileTypes()

	if !lib.MapVerifyString(fileTypes[:], strings.ToLower(f.Type)) {
		return errors.New(fmt.Sprintf("invalid type, must be: %v", fileTypes))
	}

	_, err := govalidator.ValidateStruct(f)
	if err != nil {
		return err
	}
	return nil
}

func NewFile(typ string, customer entities.ID) (*File, error) {

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

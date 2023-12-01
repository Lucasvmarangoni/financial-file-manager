package entities

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/common/const"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/common/lib"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type File struct {
	ID        string    `json:"document_id" valid:"uuid"`
	Type      string    `json:"type" valid:"notnull"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	Customer  string    `json:"customer" valid:"notnull"`
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

func NewFile(typ string, customer string) (*File, error) {

	file := File{
		ID:        uuid.NewV4().String(),
		Type:      typ,
		CreatedAt: time.Now(),
		Customer:  customer,
	}

	err := file.Validate()
	if err != nil {
		return nil, err
	}
	return &file, nil
}

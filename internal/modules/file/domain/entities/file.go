package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/const"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/lib"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/asaskevich/govalidator"
)

type File struct {
	ID         entities.ID   `json:"document_id" valid:"notnull,required"`
	Type       string        `json:"type" valid:"notnull,required"`
	CreatedAt  time.Time     `json:"created_at" valid:"-"`
	User       entities.ID   `json:"user" valid:"-"`
	Authorized []entities.ID `json:"authorized" valid:"-"`
	Versions   []entities.ID `json:"versions" valid:"-"`
	Archived   bool          `json:"archived" valid:"-"`
}

func (f *File) Validate() error {

	fileTypes := consts.FileTypes()

	if _, err := entities.ParseID(f.ID.String()); err != nil {
		errors.IsInvalidError("ID", "Must be google uuid")
	}

	if _, err := entities.ParseID(f.User.String()); err != nil {
		errors.IsInvalidError("Customer", "Must be google uuid")
	}

	if !lib.MapVerifyString(fileTypes[:], strings.ToLower(f.Type)) {
		return errors.IsInvalidError("Type", fmt.Sprintf("Must be: %v", fileTypes))
	}

	for _, versionID := range f.Versions {
		if _, err := entities.ParseID(versionID.String()); err != nil {
			return errors.IsInvalidError("Versions", "Each ID must be a google uuid")
		}
	}

	for _, authorizedID := range f.Authorized {
		if _, err := entities.ParseID(authorizedID.String()); err != nil {
			return errors.IsInvalidError("Authorized", "Each ID must be a google uuid")
		}
	}

	_, err := govalidator.ValidateStruct(f)
	if err != nil {
		return err
	}
	return nil
}

func NewFile(typ string, user entities.ID, authorized, versions []entities.ID, archived bool) (*File, error) {

	file := File{
		Type:       typ,
		User:       user,
		Authorized: authorized,
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

package entities

import (
	"errors"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/asaskevich/govalidator"
)

type Contract struct {
	File    `json:"file" valid:"required"` 
	Title   string   `json:"title" valid:"notnull"`
	Parties []string `json:"parties" valid:"notnull"`
	Object  string   `json:"object" valid:"notnull"`
	Extract []entities.ID `json:"extract" valid:"-"`
	Invoice []entities.ID `json:"invoice" valid:"-"`
}

func (c *Contract) Validate() error {

	switch {
	case c.Title == "":
		return errors.New("Need a title")
	case len(c.Parties) < 2:
		return errors.New("Insufficient number of parties")
	case c.Object == "":
		return errors.New("Need a object")
	}

	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}

func NewContract(
	file *File,
	title string,
	parties []string,
	object string,
	extract []entities.ID,
	invoice []entities.ID,
) (*Contract, error) {

	contract := &Contract{
		File:    *file,
		Title:   title,
		Parties: parties,
		Object:  object,
		Extract: extract,
		Invoice: invoice,
	}

	err := contract.Validate()
	if err != nil {
		return nil, err
	}
	return contract, nil
}

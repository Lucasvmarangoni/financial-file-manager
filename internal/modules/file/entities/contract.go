package entities

import (
	"errors"	

	"github.com/asaskevich/govalidator"
)

type Contract struct {
	File
	Title   string   `json:"title" valid:"notnull"`
	Parties []string `json:"parties" valid:"notnull"`
	Object  string   `json:"object" valid:"notnull"`
	Extract []string `json:"extract"`
	Invoice []string `json:"invoice"`
}

func (c *Contract) validate() error {	

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

func NewContract(file File, title string, parties []string, object string, extract []string, invoice []string) *Contract {
	return &Contract{
		File: file,
		Title: title,
		Parties: parties,
		Object: object,
		Extract: extract,
		Invoice: invoice,	
	}
}
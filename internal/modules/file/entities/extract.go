package entities

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type Extract struct {
	File
	Account  int     `json:"account" valid:"notnull"`
	Value    float64 `json:"value" valid:"notnull"`
	Category string  `json:"category" valid:"notnull"`
	Method   string  `json:"method" valid:"notnull"`
	Location string  `json:"location" valid:"notnull"`
	Contract *string `json:"contract,omitempty"`
}

func (e *Extract) validate() error {	

	switch {
	case e.Account == 0:
		return errors.New("Account needs to be greater than 0")
	case e.Value == 0:
		return errors.New("Value needs to be greater than 0")	
	case e.Category == "":
		return errors.New("Need a category")	
	case e.Method == "":
		return errors.New("Need a method")	
	case e.Location == "":
		return errors.New("Need a location")	
	}

	_, err := govalidator.ValidateStruct(e)

	if err != nil {
		return err
	}
	return nil
}

func NewExtract(file File, account int, value float64, category string, method string, location string, contract *string) *Extract {
    e := &Extract{
        File: file,
        Account: account,
        Value: value,
        Category: category,
        Method: method,
        Location: location,
    }
    if contract != nil {
        e.Contract = contract
    }
    return e
}
package entities

import (
	"errors"
	"fmt"
	"strings"

	consts "github.com/Lucasvmarangoni/financial-file-manager/internal/common/const"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/common/lib"
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

func (e *Extract) Validate() error {

	method := consts.Method()
	payment := consts.Payment()

	switch {
	case e.Account == 0:
		return errors.New("Account needs to be greater than 0")
	case e.Value == 0:
		return errors.New("Value needs to be greater than 0")
	case !lib.MapVerifyString(payment[:], strings.ToLower(e.Category)):
		return errors.New(fmt.Sprintf("Need a valid category: %v", payment))
	case !lib.MapVerifyString(method[:], strings.ToLower(e.Method)):
		return errors.New(fmt.Sprintf("Need a valid method: %v", method))
	case e.Location == "":
		return errors.New("Need a location")
	}

	_, err := govalidator.ValidateStruct(e)

	if err != nil {
		return err
	}
	return nil
}

func NewExtract(
	file File,
	account int,
	value float64,
	category string,
	method string,
	location string,
	contract *string,
) (*Extract, error) {

	extract := Extract{
		File:     file,
		Account:  account,
		Value:    value,
		Category: category,
		Method:   method,
		Location: location,
	}
	if contract != nil {
		extract.Contract = contract
	}

	err := extract.Validate()
	if err != nil {
		return nil, err
	}
	return &extract, nil
}

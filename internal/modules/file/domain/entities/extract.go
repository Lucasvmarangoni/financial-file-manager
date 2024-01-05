package entities

import (
	"fmt"
	"strings"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/const"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	pkg_errors "github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/lib"
	"github.com/asaskevich/govalidator"
)

type Extract struct {
	File     `json:"file" valid:"required"`
	Account  int           `json:"account" valid:"notnull"`
	Value    float64       `json:"value" valid:"notnull"`
	Category string        `json:"category" valid:"notnull,required"`
	Method   string        `json:"method" valid:"notnull,required"`
	Location string        `json:"location" valid:"required"`
	Contract entities.ID `json:"contract" valid:"-"`
}

func (e *Extract) Validate() error {

	method := consts.Method()
	payment := consts.Payment()

	if e.Account == 0 {
		return pkg_errors.IsInvalidError("Account", "Needs to be greater than 0")
	}
	if e.Value == 0 {
		return pkg_errors.IsInvalidError("Value", "Needs to be greater than 0")
	}
	if !lib.MapVerifyString(payment[:], strings.ToLower(e.Category)) {
		return pkg_errors.IsInvalidError("Category", fmt.Sprintf("It need to be a one of those: %v", payment))
	}
	if !lib.MapVerifyString(method[:], strings.ToLower(e.Method)) {
		return pkg_errors.IsInvalidError("Method", fmt.Sprintf("It need to be a one of those: %v", method))
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
	contract entities.ID,
) (*Extract, error) {

	extract := Extract{
		File:     file,
		Account:  account,
		Value:    value,
		Category: category,
		Method:   method,
		Location: location,
	}

	err := extract.Validate()
	if err != nil {
		return nil, err
	}
	return &extract, nil
}

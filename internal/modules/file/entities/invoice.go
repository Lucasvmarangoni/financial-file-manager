package entities

import (
	"errors"
	"fmt"
	"time"

	consts "github.com/Lucasvmarangoni/financial-file-manager/internal/common/const"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/common/lib"
	"github.com/asaskevich/govalidator"
)

type Invoice struct {
	File
	DueDate  time.Time `json:"maturity" valid:"notnull"`
	Value    float64   `json:"value" valid:"notnull"`
	Method   string    `json:"method" valid:"notnull"`
	Contract *string   `json:"contract,omitempty"`
}

func (i *Invoice) Validate() error {

	methods := consts.Method()

	switch {
	case i.Value <= 0:
		return errors.New("Value needs to be greater than 0")
	case i.DueDate == time.Time{}:
		return errors.New("Need a due date")
	case !lib.MapVerifyString(methods[:], i.Method):
		return errors.New(fmt.Sprintf("Need a valid method: %v", methods))
	}

	_, err := govalidator.ValidateStruct(i)

	if err != nil {
		return err
	}
	return nil
}

func NewInvoice(
	file File,
	dueDate time.Time,
	value float64,
	method string,
	contract *string,
) (*Invoice, error) {
	invoice := &Invoice{
		File:    file,
		DueDate: dueDate,
		Value:   value,
		Method:  method,
	}
	if contract != nil {
		invoice.Contract = contract
	}
	err := invoice.Validate()
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

package entities

import (
	"errors"
	"time"

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
	switch {
	case i.Value <= 0:
		return errors.New("Value needs to be greater than 0")
	case i.DueDate == time.Time{}:
		return errors.New("Need a due date")
	case i.Method == "":
		return errors.New("Need a method")
	}

	_, err := govalidator.ValidateStruct(i)

	if err != nil {
		return err
	}
	return nil
}

func NewInvoice(file File, dueDate time.Time, value float64, method string, contract *string) *Invoice {
	i := &Invoice{
		File:    file,
		DueDate: dueDate,
		Value:   value,
		Method:  method,
	}
	if contract != nil {
		i.Contract = contract
	}
	return i
}

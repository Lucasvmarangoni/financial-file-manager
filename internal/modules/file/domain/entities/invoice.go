package entities

import (
	"fmt"
	"time"

	"github.com/Lucasvmarangoni/logella/err"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/const"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/lib"
	"github.com/asaskevich/govalidator"
)

type Invoice struct {
	File     `json:"file" valid:"required"`
	DueDate  time.Time   `json:"maturity" valid:"notnull"`
	Value    float64     `json:"value" valid:"notnull"`
	Method   string      `json:"method" valid:"notnull,required"`
	Contract entities.ID `json:"contract" valid:"-"`
}

func (i *Invoice) Validate() error {

	method := consts.Method()

	if i.Value <= 0 {
		return errors.IsRequiredError("Value", "It needs to be greater than 0")
	}
	if i.DueDate.IsZero() {
		return errors.IsRequiredError("DueDate", "")
	}
	if !lib.MapVerifyString(method[:], i.Method) {
		return errors.IsInvalidError("Method", fmt.Sprintf("It need to be a one of those: %v", method))
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
	contract entities.ID,
) (*Invoice, error) {
	invoice := &Invoice{
		File:    file,
		DueDate: dueDate,
		Value:   value,
		Method:  method,
	}

	err := invoice.Validate()
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

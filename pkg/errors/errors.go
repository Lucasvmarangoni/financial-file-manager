package errors

import (
	"errors"
	"fmt"
)

func IsRequiredError(fieldName, msg string) error {
	return errors.New(fmt.Sprintf("%s is required. %s", fieldName, msg))
}

func IsInvalidError(fieldName, msg string) error {
	return errors.New(fmt.Sprintf("%s is invalid. %s", fieldName, msg))
}

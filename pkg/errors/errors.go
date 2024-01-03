package errors

import (
	"errors"
	"fmt"
)

type CustomError struct {
	err   error
	added bool
}

var e *CustomError

func NewError(err error, value ...string) error {
	key := "Operation"

	if e == nil || !e.added {
		e = &CustomError{
			err:   fmt.Errorf("%s: %w", "Error", err),
			added: true,
		}
	} else {
		e = &CustomError{
			err: err,
		}
	}
	return fmt.Errorf("%s: %s - %w", key, value, e.err)
}

func IsRequiredError(fieldName, msg string) error {
	return errors.New(fmt.Sprintf("%s is required. %s", fieldName, msg))
}

func IsInvalidError(fieldName, msg string) error {
	return errors.New(fmt.Sprintf("%s is invalid. %s", fieldName, msg))
}

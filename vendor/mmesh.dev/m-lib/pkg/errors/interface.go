package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

func New(msg string) error {
	return fmt.Errorf(msg)
}

func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

func Is(err, targetErr error) bool {
	return errors.Is(err, targetErr)
}

func Cause(err error) error {
	return errors.Cause(err)
}

func Errs(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	errMsg := "ERRORS: |+| "
	for _, err := range errs {
		errMsg = errMsg + fmt.Sprint(err) + " |+| "
	}
	return New(errMsg)
}

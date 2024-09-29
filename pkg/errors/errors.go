package errors

import "github.com/pkg/errors"

func New(msg string) error {
	return errors.New(msg)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func AggregateErrors(errs []error) error {
	var result error
	for _, err := range errs {
		if result == nil {
			result = err
		} else {
			result = errors.Wrap(result, err.Error())
		}
	}
	return result
}

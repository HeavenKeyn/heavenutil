package errutil

import (
	"errors"
	"fmt"
)

type MultiErrors struct {
	errors []error
}

func (e *MultiErrors) Append(err error) {
	if e.errors == nil {
		e.errors = make([]error, 0)
	}
	if err != nil {
		e.errors = append(e.errors, err)
	}
}

func (e *MultiErrors) AddError(args ...interface{}) {
	e.Append(Error(args...))
}

func (e *MultiErrors) GetErrors() []error {
	return e.errors
}

func (e *MultiErrors) GetError() error {
	if e.errors == nil || len(e.errors) == 0 {
		return nil
	}
	return errors.New(fmt.Sprint(e.errors))
}

func Error(args ...interface{}) error {
	return errors.New(fmt.Sprint(args...))
}

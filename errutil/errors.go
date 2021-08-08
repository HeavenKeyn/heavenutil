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
	e.errors = append(e.errors, err)
}

func (e *MultiErrors) AddError(args ...interface{}) {
	if e.errors == nil {
		e.errors = make([]error, 0)
	}
	e.errors = append(e.errors, Error(args...))
}

func (e *MultiErrors) GetErrors() []error {
	return e.errors
}

func (e *MultiErrors) GetError() error {
	return errors.New(fmt.Sprint(e.errors))
}

func Error(args ...interface{}) error {
	return errors.New(fmt.Sprint(args...))
}

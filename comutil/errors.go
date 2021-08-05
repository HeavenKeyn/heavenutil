package comutil

import (
	"errors"
	"fmt"
)

type MultiErrors struct {
	errors []error
}

func (e *MultiErrors) AddError(err error) {
	if e.errors == nil {
		e.errors = make([]error, 0)
	}
	e.errors = append(e.errors, err)
}

func (e *MultiErrors) GetErrors() []error {
	return e.errors
}

func (e *MultiErrors) GetError() error {
	return errors.New(fmt.Sprint(e.errors))
}

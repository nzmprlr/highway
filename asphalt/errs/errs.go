package errs

import (
	"errors"
)

const (
	emptyErrorMsg = "empty error message"
)

type Errs struct {
	errors []error
}

func (e *Errs) String() (s string) {
	if e == nil || len(e.errors) == 0 {
		return
	}

	for _, err := range e.errors {
		errMsg := err.Error()
		if errMsg == "" {
			errMsg = emptyErrorMsg
		}

		l := len(errMsg)
		if l > 0 && errMsg[:l-1] != "\n" {
			errMsg += "\n"
		}

		s += errMsg
	}

	// to avoid panic of s[:l-1].
	l := len(s)
	if l == 0 {
		return
	}

	return s[:l-1]
}

func (e *Errs) Error() string {
	return e.String()
}

func (e *Errs) Reduce() error {
	if e == nil || len(e.errors) == 0 {
		return nil
	}

	return e
}

func (e *Errs) Add(err error) error {
	if e == nil {
		if err != nil && err.Error() == "" {
			return errors.New(emptyErrorMsg)
		}

		return err
	}

	if err == nil {
		return e
	}

	if e.errors == nil {
		e.errors = []error{}
	}

	e.errors = append(e.errors, err)
	return e
}

func New() *Errs {
	return &Errs{
		errors: []error{},
	}
}

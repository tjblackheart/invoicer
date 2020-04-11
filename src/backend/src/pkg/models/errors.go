package models

import "fmt"

type (
	ErrUnique struct {
		Message string
	}

	ErrInvalidCredentials struct {
		Message string
	}

	ErrNotFound struct {
		Message string
	}

	ErrValidation struct {
		Message string
	}
)

func (e ErrUnique) Error() string {
	return fmt.Sprintf(e.Message)
}

func (e ErrInvalidCredentials) Error() string {
	return fmt.Sprintf(e.Message)
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf(e.Message)
}

func (e ErrValidation) Error() string {
	return fmt.Sprintf(e.Message)
}

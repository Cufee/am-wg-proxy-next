package errors

import "fmt"

const (
	TypeInput = iota
	TypeGeneric
	TypeInternal
)

type Error struct {
	Raw     error  `json:"-"`
	Type    int    `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%v] %s: %s", e.Type, e.Message, e.Raw.Error())
}

func Generic(err error, message string) *Error {
	return &Error{
		Raw:     err,
		Type:    TypeGeneric,
		Message: message,
	}
}

func Internal(err error, message string) *Error {
	return &Error{
		Raw:     err,
		Type:    TypeInternal,
		Message: message,
	}
}

func Input(err error, message string) *Error {
	return &Error{
		Raw:     err,
		Type:    TypeInput,
		Message: message,
	}
}

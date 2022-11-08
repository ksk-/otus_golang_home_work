package hw09structvalidator

import "errors"

var (
	ErrNoStruct                = errors.New("value isn't struct")
	ErrInvalidValidationRule   = errors.New("invalid validation rule")
	ErrUnknownValidationRule   = errors.New("unknown validation rule")
	ErrDuplicateValidationRule = errors.New("duplicate validation rule")
	ErrTooShortString          = errors.New("too short string")
	ErrTooLongString           = errors.New("too long string")
	ErrInvalidPattern          = errors.New("invalid pattern")
	ErrUnknownValue            = errors.New("unknown value")
	ErrTooSmallNumber          = errors.New("too small number")
	ErrTooBigNumber            = errors.New("too big number")
)

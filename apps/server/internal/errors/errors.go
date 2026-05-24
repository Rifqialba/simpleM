package errors

import "errors"

var (
	ErrInvalidRequest = errors.New("invalid request")

	ErrValidation = errors.New("validation failed")

	ErrInternalServer = errors.New("internal server error")
)
package errors

import "errors"

var (
	ErrServiceNotFound = errors.New("service not found")
	ErrInvalidArgument = errors.New("invalid argument")
)

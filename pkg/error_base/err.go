package errorbase

import "errors"

var (
	ErrInvalidId              = errors.New("invalid id")
	ErrEmptyParameters        = errors.New("empty parameters")
	ErrEmptyList              = errors.New("empty list")
	ErrConflict               = errors.New("element already exist")
	ErrNotFound               = errors.New("not found")
	ErrModelInvalid           = errors.New("model invalid")
	ErrStorageOperationFailed = errors.New("storage operation failed")
	ErrInvalidNumber          = errors.New("invalid number")
	ErrInvalidRequest         = errors.New("invalid request")
	ErrUnprocessable          = errors.New("incorrect request")
)

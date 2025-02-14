package errorbase

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidId               = errors.New("invalid id")
	ErrEmptyParameters         = errors.New("empty parameters")
	ErrEmptyList               = errors.New("empty list")
	ErrConflict                = errors.New("element already exist")
	ErrNotFound                = errors.New("not found")
	ErrModelInvalid            = errors.New("model invalid")
	ErrStorageOperationFailed  = errors.New("storage operation failed")
	ErrDatabaseOperationFailed = errors.New("database operation failed")
	ErrInvalidNumber           = errors.New("invalid number")
	ErrInvalidRequest          = errors.New("invalid request")
	ErrUnprocessable           = errors.New("incorrect request")
	ErrBuyerFKNotExist         = errors.New("buyer FK not exist")
	ErrCarrierFKNotExist       = errors.New("carrier FK not exist")
	ErrOrderStatusFKNotExist   = errors.New("order status FK not exist")
	ErrWareHouseFKNotExist     = errors.New("warehouse FK  not exist")
	ErrOrderNumberExist        = errors.New("order number already exist")
	ErrTrackingCodeExist       = errors.New("tracking code already exist")
	ErrInvalidIdField          = errors.New("invalid id field")
)

type ErrorFKNotExist struct {
	FK []string
}

func (e *ErrorFKNotExist) Error() string {
	return fmt.Sprintf("FK invalid: %s", strings.Join(e.FK, ", "))
}

func (e *ErrorFKNotExist) AddFK(fk string) {
	e.FK = append(e.FK, fk)
}

func (e *ErrorFKNotExist) HasErrors() bool {
	return len(e.FK) > 0
}

func (e *ErrorFKNotExist) Is(target error) bool {
	_, ok := target.(*ErrorFKNotExist)
	return ok
}

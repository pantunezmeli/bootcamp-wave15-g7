package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInsertingData          = errors.New("error inserting data on database")
	ErrGettingLastID          = errors.New("error getting last insert ID")
	ErrConvertingID           = errors.New("error parsing ID")
	ErrForeignKeyViolation    = errors.New("error invalid foreign key")
	ErrWarehouseCodeDuplicate = errors.New("error warehouse code already exists")
	ErrDuplicateEntry         = errors.New("error cid already exists")
	ErrDBGenericError         = errors.New("error database failed")
	ErrLocalityNotFound       = errors.New("locality does not exists")
	ErrMappingData            = errors.New("error mapping data on scan")
	ErrWarehouseNotFound      = errors.New("warehouse not found")
)

// Validations
type ErrInvalidParameter struct {
	Parameter string
}

func (e ErrInvalidParameter) Error() string {
	return fmt.Sprintf("missing parameter: %s", e.Parameter)
}

// Database
type ErrDatabase struct {
	Err error
}

func (e ErrDatabase) Error() string {
	return fmt.Sprintf("database error: %s", e.Err)
}

func (e ErrDatabase) Unwrap() error {
	return e.Err
}

// Foreign Key
type ErrForeignKey struct {
	Err error
}

func (e ErrForeignKey) Error() string {
	return fmt.Sprintf("foreign key violation: %v", e.Err)
}

func (e ErrForeignKey) Unwrap() error {
	return e.Err
}

// Duplicate
type ErrDuplicate struct {
	Err error
}

func (e ErrDuplicate) Error() string {
	return fmt.Sprintf("duplicate entry: %v", e.Err)
}

func (e ErrDuplicate) Unwrap() error {
	return e.Err
}

// Convertion
type ErrConvertion struct {
	Err error
}

func (e ErrConvertion) Error() string {
	return fmt.Sprintf("failed to convert to DTO: %v", e.Err)
}

func (e ErrConvertion) Unwrap() error {
	return e.Err
}

// Not Found
type ErrNotFound struct {
	Err error
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("element not found: %v", e.Err)
}

func (e ErrNotFound) Unwrap() error {
	return e.Err
}

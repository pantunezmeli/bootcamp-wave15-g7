package carrier

import (
	"fmt"

	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/carrier"
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

type ICarrierService interface {

	// ! 1)
	AddCarrier(req dto.CarrierDoc) (c dto.CarrierDoc, err error)
}

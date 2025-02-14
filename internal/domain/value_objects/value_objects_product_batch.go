package value_objects

import (
	"database/sql/driver"
	"fmt"
	"time"

	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type BatchNumber string

type CurrentQuantity int

type DueDate time.Time

func (d DueDate) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d *DueDate) Scan(value interface{}) error {
	if value == nil {
		*d = DueDate(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = DueDate(v)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into DueDate", value)
	}
}

type InitialQuantity int

type ManufacturingDate time.Time

func (d ManufacturingDate) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d *ManufacturingDate) Scan(value interface{}) error {
	if value == nil {
		*d = ManufacturingDate(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = ManufacturingDate(v)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into ManufacturingDate", value)
	}
}

type ManufacturingHour time.Time

func (d ManufacturingHour) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d *ManufacturingHour) Scan(value interface{}) error {
	if value == nil {
		*d = ManufacturingHour(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = ManufacturingHour(v)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into ManufacturingHour", value)
	}
}

// NewBatchNumber is a function that creates a new batch number
func NewBatchNumber(batch_number string) (BatchNumber, error) {
	if batch_number == "" {
		return "", errorbase.ErrInvalidBatchNumber
	}

	return BatchNumber(batch_number), nil
}

// NewCurrentQuantity is a function that creates a new current quantity
func NewCurrentQuantity(current_quantity int) (CurrentQuantity, error) {
	if current_quantity == 0 {
		return 0, errorbase.ErrInvalidCurrentQuantity
	}

	return CurrentQuantity(current_quantity), nil
}

// NewDueDate is a function that creates a new due date
func NewDueDate(due_date time.Time) (DueDate, error) {
	if due_date.IsZero() {
		return DueDate{}, errorbase.ErrInvalidDueDate
	}

	return DueDate(due_date), nil
}

// NewInitialQuantity is a function that creates a new initial quantity
func NewInitialQuantity(initial_quantity int) (InitialQuantity, error) {
	if initial_quantity == 0 {
		return 0, errorbase.ErrInvalidInitialQuantity
	}

	return InitialQuantity(initial_quantity), nil
}

// NewManufacturingDate is a function that creates a new manufacturing date
func NewManufacturingDate(manufacturing_date time.Time) (ManufacturingDate, error) {
	if manufacturing_date.IsZero() {
		return ManufacturingDate{}, errorbase.ErrInvalidManufacturingDate
	}

	return ManufacturingDate(manufacturing_date), nil
}

// NewManufacturingHour is a function that creates a new manufacturing hour
func NewManufacturingHour(manufacturing_hour time.Time) (ManufacturingHour, error) {
	if manufacturing_hour.IsZero() {
		return ManufacturingHour{}, errorbase.ErrInvalidManufacturingHour
	}

	return ManufacturingHour(manufacturing_hour), nil
}

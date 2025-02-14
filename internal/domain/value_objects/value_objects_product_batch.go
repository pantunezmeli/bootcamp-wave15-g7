package value_objects

import (
	"fmt"
	"time"

	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type ProductBatchId int

type BatchNumber string

type CurrentQuantity int

type DueDate time.Time

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

type ManufacturingHour time.Time

func NewProductBatchId(id int) (ProductBatchId, error) {
	if id == 0 {
		return 0, errorbase.ErrInvalidId
	}

	return ProductBatchId(id), nil
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
func NewDueDate(due_date string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", due_date)
	if err != nil || parsedDate.Year() < 1 || parsedDate.Year() > 9999 {
		return time.Time{}, errorbase.ErrInvalidDueDate
	}
	return parsedDate, nil
}

// NewInitialQuantity is a function that creates a new initial quantity
func NewInitialQuantity(initial_quantity int) (InitialQuantity, error) {
	if initial_quantity == 0 {
		return 0, errorbase.ErrInvalidInitialQuantity
	}

	return InitialQuantity(initial_quantity), nil
}

// NewManufacturingDate is a function that creates a new manufacturing date
func NewManufacturingDate(manufacturing_date string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", manufacturing_date)
	if err != nil || parsedDate.Year() < 1 || parsedDate.Year() > 9999 {
		return time.Time{}, errorbase.ErrInvalidDueDate
	}

	return parsedDate, nil
}

// NewManufacturingHour is a function that creates a new manufacturing hour
func NewManufacturingHour(manufacturing_hour string) (time.Time, error) {
	parsedHour, err := time.Parse("15:04:05", manufacturing_hour)
	if err != nil {
		return time.Time{}, errorbase.ErrInvalidManufacturingHour
	}

	// Set the date to a default value, e.g., 0000-01-01
	defaultDate := time.Date(1, 1, 1, parsedHour.Hour(), parsedHour.Minute(), parsedHour.Second(), 0, time.UTC)

	return defaultDate, nil
}

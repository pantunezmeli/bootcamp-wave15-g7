package value_objects

import (
	"time"

	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type BatchNumber struct {
	batch_number string
}

type CurrentQuantity struct {
	current_quantity int
}

type DueDate struct {
	due_date time.Time
}

type InitialQuantity struct {
	initial_quantity int
}

type ManufacturingDate struct {
	manufacturing_date time.Time
}

type ManufacturingHour struct {
	manufacturing_hour time.Time
}

type MinumumTemperature struct {
	minumum_temperature float64
}

// GetBatchNumber is a function that returns the batch number of a product batch
func (b BatchNumber) GetBatchNumber() string {
	return b.batch_number
}

// GetCurrentQuantity is a function that returns the current quantity of a product batch
func (c CurrentQuantity) GetCurrentQuantity() int {
	return c.current_quantity
}

func (d DueDate) GetDueDate() time.Time {
	return d.due_date
}

func (i InitialQuantity) GetInitialQuantity() int {
	return i.initial_quantity
}

func (m ManufacturingDate) GetManufacturingDate() time.Time {
	return m.manufacturing_date
}

func (m ManufacturingHour) GetManufacturingHour() time.Time {
	return m.manufacturing_hour
}

func (m MinumumTemperature) GetMinumumTemperature() float64 {
	return m.minumum_temperature
}

// NewBatchNumber is a function that creates a new batch number
func NewBatchNumber(batch_number string) (BatchNumber, error) {
	if batch_number == "" {
		return BatchNumber{}, errorbase.ErrInvalidBatchNumber
	}

	return BatchNumber{
		batch_number: batch_number,
	}, nil
}

// NewCurrentQuantity is a function that creates a new current quantity
func NewCurrentQuantity(current_quantity int) (CurrentQuantity, error) {
	if current_quantity == 0 {
		return CurrentQuantity{}, errorbase.ErrInvalidCurrentQuantity
	}

	return CurrentQuantity{
		current_quantity: current_quantity,
	}, nil
}

// NewDueDate is a function that creates a new due date
func NewDueDate(due_date time.Time) (DueDate, error) {
	if due_date.IsZero() {
		return DueDate{}, errorbase.ErrInvalidDueDate
	}

	return DueDate{
		due_date: due_date,
	}, nil
}

// NewInitialQuantity is a function that creates a new initial quantity
func NewInitialQuantity(initial_quantity int) (InitialQuantity, error) {
	if initial_quantity == 0 {
		return InitialQuantity{}, errorbase.ErrInvalidInitialQuantity
	}

	return InitialQuantity{
		initial_quantity: initial_quantity,
	}, nil
}

// NewManufacturingDate is a function that creates a new manufacturing date
func NewManufacturingDate(manufacturing_date time.Time) (ManufacturingDate, error) {
	if manufacturing_date.IsZero() {
		return ManufacturingDate{}, errorbase.ErrInvalidManufacturingDate
	}

	return ManufacturingDate{
		manufacturing_date: manufacturing_date,
	}, nil
}

// NewManufacturingHour is a function that creates a new manufacturing hour
func NewManufacturingHour(manufacturing_hour time.Time) (ManufacturingHour, error) {
	if manufacturing_hour.IsZero() {
		return ManufacturingHour{}, errorbase.ErrInvalidManufacturingHour
	}

	return ManufacturingHour{
		manufacturing_hour: manufacturing_hour,
	}, nil
}

// NewMinumumTemperature is a function that creates a new minumum temperature
func NewMinumumTemperature(minumum_temperature float64) (MinumumTemperature, error) {
	if minumum_temperature == 0 {
		return MinumumTemperature{}, errorbase.ErrInvalidMinumumTemperature
	}

	return MinumumTemperature{
		minumum_temperature: minumum_temperature,
	}, nil
}

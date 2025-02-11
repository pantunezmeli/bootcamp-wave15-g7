package models

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"

// ProductBatch is a struct that represents a product batch
type ProductBatch struct {
	ID                 value_objects.Id
	BatchNumber        int
	CurrentQuantity    int
	CurrentTemperature float64
	DueDate            string
	InitialQuantity    int
	ManufacturingDate  string
	ManufacturingHour  string
	MinumumTemperature float64
	ProductID          value_objects.Id
	SectionID          value_objects.Id
}

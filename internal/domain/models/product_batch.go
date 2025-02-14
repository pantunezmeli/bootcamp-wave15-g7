package models

import (
	"time"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

// ProductBatch is a struct that represents a product batch
type ProductBatch struct {
	Id                 value_objects.ProductBatchId
	BatchNumber        value_objects.BatchNumber
	CurrentQuantity    value_objects.CurrentQuantity
	CurrentTemperature value_objects.CurrentTemperature
	DueDate            time.Time
	InitialQuantity    value_objects.InitialQuantity
	ManufacturingDate  time.Time
	ManufacturingHour  time.Time
	MinimumTemperature value_objects.MinimumTemperature
	ProductID          value_objects.Id
	SectionID          value_objects.Id
}

// // NewProductBatch is a function that creates a new product batch
// func NewProductBatch(id value_objects.Id, batchNumber string, currentQuantity int, currentTemperature float64, dueDate time.Time, initialQuantity int, manufacturingDate time.Time, manufacturingHour time.Time, minimumTemperature float64, productID value_objects.Id, sectionID value_objects.Id) (ProductBatch, error) {
// 	if id.GetId() == 0 {
// 		return ProductBatch{}, errorbase.ErrInvalidId
// 	}
// 	if batchNumber == "" {
// 		return ProductBatch{}, errorbase.ErrInvalidBatchNumber
// 	}
// 	if currentQuantity == 0 {
// 		return ProductBatch{}, errorbase.ErrInvalidCurrentQuantity
// 	}
// 	if currentTemperature == 0 {
// 		return ProductBatch{}, errorbase.ErrInvalidCurrentTemperature
// 	}
// 	if dueDate.IsZero() {
// 		return ProductBatch{}, errorbase.ErrInvalidDueDate
// 	}
// 	if initialQuantity == 0 {
// 		return ProductBatch{}, errorbase.ErrInvalidInitialQuantity
// 	}
// 	if manufacturingDate.IsZero() {
// 		return ProductBatch{}, errorbase.ErrInvalidManufacturingDate
// 	}
// 	if manufacturingHour.IsZero() {
// 		return ProductBatch{}, errorbase.ErrInvalidManufacturingHour
// 	}
// 	if minimumTemperature == 0 {
// 		return ProductBatch{}, errorbase.ErrInvalidMinumumTemperature
// 	}
// 	if productID.GetId() == 0 {
// 		return ProductBatch{}, errorbase.ErrInvalidProductID
// 	}
// 	if sectionID.GetId() == 0 {
// 		return ProductBatch{}, errorbase.ErrInvalidSectionID
// 	}

// 	BatchNumber, err := value_objects.NewBatchNumber(batchNumber)
// 	if err != nil {
// 		return ProductBatch{}, err
// 	}
// 	CurrentQuantity, err := value_objects.NewCurrentQuantity(currentQuantity)
// 	if err != nil {
// 		return ProductBatch{}, err
// 	}
// 	CurrentTemperature, err := value_objects.NewCurrentTemperature(currentTemperature)
// 	if err != nil {
// 		return ProductBatch{}, err
// 	}
// 	DueDate, err := value_objects.NewDueDate(dueDate)
// 	if err != nil {
// 		return ProductBatch{}, err
// 	}
// 	InitialQuantity, err := value_objects.NewInitialQuantity(initialQuantity)
// 	if err != nil {
// 		return ProductBatch{}, err
// 	}
// 	ManufacturingDate, err := value_objects.NewManufacturingDate(manufacturingDate)
// 	if err != nil {
// 		return ProductBatch{}, err
// 	}
// 	ManufacturingHour, err := value_objects.NewManufacturingHour(manufacturingHour)
// 	if err != nil {
// 		return ProductBatch{}, err
// 	}
// 	MinimumTemperature, err := value_objects.NewMinimumTemperature(minimumTemperature)
// 	if err != nil {
// 		return ProductBatch{}, err
// 	}

// 	return ProductBatch{
// 		Id:                 id,
// 		BatchNumber:        BatchNumber,
// 		CurrentQuantity:    CurrentQuantity,
// 		CurrentTemperature: CurrentTemperature,
// 		DueDate:            DueDate,
// 		InitialQuantity:    InitialQuantity,
// 		ManufacturingDate:  ManufacturingDate,
// 		ManufacturingHour:  ManufacturingHour,
// 		MinimumTemperature: MinimumTemperature,
// 		ProductID:          productID,
// 		SectionID:          sectionID,
// 	}, nil
// }

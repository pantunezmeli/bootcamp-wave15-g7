package value_objects

import (
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type SectionNumber struct {
	section_Number int
}

type CurrentTemperature struct {
	current_Temperature float64
}

type MinimumTemperature struct {
	minimum_Temperature int
}

type CurrentCapacity struct {
	current_Capacity int
}

type MinimumCapacity struct {
	minimum_Capacity int
}

type MaximumCapacity struct {
	maximum_Capacity int
}

type WarehouseId struct {
	warehouse_Id int
}

type ProductTypeId struct {
	product_Type_Id int
}

func NewSectionNumber(sectionNumber int) (SectionNumber, error) {
	if sectionNumber <= 0 {
		return SectionNumber{}, errorbase.ErrInvalidId
	}
	return SectionNumber{section_Number: sectionNumber}, nil
}

func (f SectionNumber) GetSectionNumber() int {
	return f.section_Number
}

func NewCurrentTemperature(currentTemperature float64) (CurrentTemperature, error) {
	if currentTemperature <= 0 {
		return CurrentTemperature{}, errorbase.ErrInvalidNumber
	}
	return CurrentTemperature{current_Temperature: currentTemperature}, nil
}

func (l CurrentTemperature) GetCurrentTemperature() float64 {
	return l.current_Temperature
}

func NewMinimumTemperature(minimumTemperature int) (MinimumTemperature, error) {
	if minimumTemperature <= 0 {
		return MinimumTemperature{}, errorbase.ErrInvalidNumber
	}
	return MinimumTemperature{minimum_Temperature: minimumTemperature}, nil
}

func (l MinimumTemperature) GetMinimumTemperature() int {
	return l.minimum_Temperature
}

func NewCurrentCapacity(currentCapacity int) (CurrentCapacity, error) {
	if currentCapacity < 0 || currentCapacity > 1000 {
		return CurrentCapacity{}, errorbase.ErrInvalidId
	}
	return CurrentCapacity{current_Capacity: currentCapacity}, nil
}

func (l CurrentCapacity) GetCurrentCapacity() int {
	return l.current_Capacity
}

func NewMinimumCapacity(minimumCapacity int) (MinimumCapacity, error) {
	if minimumCapacity < 0 || minimumCapacity > 1000 {
		return MinimumCapacity{}, errorbase.ErrInvalidId
	}
	return MinimumCapacity{minimum_Capacity: minimumCapacity}, nil
}

func (l MinimumCapacity) GetMinimumCapacity() int {
	return l.minimum_Capacity
}

func NewMaximumCapacity(maximumCapacity int) (MaximumCapacity, error) {
	if maximumCapacity < 0 || maximumCapacity > 2000 {
		return MaximumCapacity{}, errorbase.ErrInvalidId
	}
	return MaximumCapacity{maximum_Capacity: maximumCapacity}, nil
}

func (l MaximumCapacity) GetMaximumCapacity() int {
	return l.maximum_Capacity
}

func NewWarehouseId(warehouseId int) (WarehouseId, error) {
	if warehouseId <= 0 {
		return WarehouseId{}, errorbase.ErrInvalidId
	}
	return WarehouseId{warehouse_Id: warehouseId}, nil
}

func (l WarehouseId) GetWarehouseId() int {
	return l.warehouse_Id
}

func NewProductTypeId(productTypeId int) (ProductTypeId, error) {
	if productTypeId <= 0 {
		return ProductTypeId{}, errorbase.ErrInvalidId
	}
	return ProductTypeId{product_Type_Id: productTypeId}, nil
}

func (l ProductTypeId) GetProductTypeId() int {
	return l.product_Type_Id
}

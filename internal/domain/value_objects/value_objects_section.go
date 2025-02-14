package value_objects

import (
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type SectionId int

type SectionNumber int

type CurrentTemperature float64

type MinimumTemperature float64

type CurrentCapacity int

type MinimumCapacity int

type MaximumCapacity int

type WarehouseId int

type ProductTypeId int

func NewSectionId(sectionId int) (SectionId, error) {
	if sectionId <= 0 {
		return 0, errorbase.ErrInvalidId
	}
	return SectionId(sectionId), nil
}

func NewSectionNumber(sectionNumber int) (SectionNumber, error) {
	if sectionNumber <= 0 {
		return 0, errorbase.ErrInvalidId
	}
	return SectionNumber(sectionNumber), nil
}

func NewCurrentTemperature(currentTemperature float64) (CurrentTemperature, error) {
	if currentTemperature <= 0 {
		return 0, errorbase.ErrInvalidNumber
	}
	return CurrentTemperature(currentTemperature), nil
}

func NewMinimumTemperature(minimumTemperature int) (MinimumTemperature, error) {
	if minimumTemperature <= 0 {
		return 0, errorbase.ErrInvalidNumber
	}
	return MinimumTemperature(minimumTemperature), nil
}

func NewCurrentCapacity(currentCapacity int) (CurrentCapacity, error) {
	if currentCapacity < 0 || currentCapacity > 1000 {
		return 0, errorbase.ErrInvalidId
	}
	return CurrentCapacity(currentCapacity), nil
}

func NewMinimumCapacity(minimumCapacity int) (MinimumCapacity, error) {
	if minimumCapacity < 0 || minimumCapacity > 1000 {
		return 0, errorbase.ErrInvalidId
	}
	return MinimumCapacity(minimumCapacity), nil
}

func NewMaximumCapacity(maximumCapacity int) (MaximumCapacity, error) {
	if maximumCapacity < 0 || maximumCapacity > 2000 {
		return 0, errorbase.ErrInvalidId
	}
	return MaximumCapacity(maximumCapacity), nil
}

func NewWarehouseId(warehouseId int) (WarehouseId, error) {
	if warehouseId <= 0 {
		return 0, errorbase.ErrInvalidId
	}
	return WarehouseId(warehouseId), nil
}

func NewProductTypeId(productTypeId int) (ProductTypeId, error) {
	if productTypeId <= 0 {
		return 0, errorbase.ErrInvalidId
	}
	return ProductTypeId(productTypeId), nil
}

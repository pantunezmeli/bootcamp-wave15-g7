package models

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"


type Product struct {
	ID            value_objects.Id
	ProductCode   string
	Description   string
	ProductTypeID value_objects.Id
	SellerID      value_objects.Id
	Dimensions
	FreezingInfo
}

type Dimensions struct {
	Height    float64
	Length    float64
	NetWeight float64
	Width     float64
}

type FreezingInfo struct {
	ExpirationRate                 float64
	FreezingRate                   float64
	RecommendedFreezingTemperature float64
}

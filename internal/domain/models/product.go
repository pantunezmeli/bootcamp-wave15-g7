package models

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"

type Product struct {
	ID            domain.Id
	ProductCode   string
	Description   string
	ProductTypeID domain.Id
	SellerID      domain.Id
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

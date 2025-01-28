package model

type Product struct {
	ID            int
	ProductCode   string
	Description   string
	ProductTypeID int
	SellerID      int
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

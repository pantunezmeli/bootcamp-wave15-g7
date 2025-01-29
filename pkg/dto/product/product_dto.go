package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"sort"
)

type ProductDTO struct {
	ID                             int     `json:"id,omitempty"`
	Description                    string  `json:"description,omitempty"`
	ExpirationRate                 float64 `json:"expiration_rate,omitempty"`
	FreezingRate                   float64 `json:"freezing_rate,omitempty"`
	Height                         float64 `json:"height,omitempty"`
	Length                         float64 `json:"length,omitempty"`
	NetWeight                      float64 `json:"netweight,omitempty"`
	ProductCode                    string  `json:"product_code,omitempty"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature,omitempty"`
	Width                          float64 `json:"width,omitempty"`
	ProductTypeID                  int     `json:"product_type_id,omitempty"`
	SellerID                       int     `json:"seller_id,omitempty"`
}

func ParseDTOProduct(productTypeID, sellerID domain.Id, product ProductDTO) models.Product {
	id, _ := domain.NewId(product.ID)
	return models.Product{
		ID:            id,
		Description:   product.Description,
		ProductCode:   product.ProductCode,
		ProductTypeID: productTypeID,
		SellerID:      sellerID,
		FreezingInfo: models.FreezingInfo{
			ExpirationRate:                 product.ExpirationRate,
			FreezingRate:                   product.FreezingRate,
			RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		},
		Dimensions: models.Dimensions{
			Height:    product.Height,
			Length:    product.Length,
			NetWeight: product.NetWeight,
			Width:     product.Width,
		},
	}
}

func ParserMapProductToListDTO(p map[int]models.Product) []ProductDTO {
	list := make([]ProductDTO, 0, len(p))
	for _, v := range p {
		list = append(list, ParserProductToDTO(v))
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}

func ParserProductToDTO(p models.Product) ProductDTO {
	return ProductDTO{
		ID:                             p.ID.Value(),
		Description:                    p.Description,
		ExpirationRate:                 p.ExpirationRate,
		FreezingRate:                   p.FreezingRate,
		Height:                         p.Height,
		Length:                         p.Length,
		NetWeight:                      p.NetWeight,
		ProductCode:                    p.ProductCode,
		RecommendedFreezingTemperature: p.RecommendedFreezingTemperature,
		Width:                          p.Width,
		ProductTypeID:                  p.ProductTypeID.Value(),
		SellerID:                       p.SellerID.Value(),
	}
}

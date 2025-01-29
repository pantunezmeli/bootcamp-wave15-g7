package dto

import (
	"errors"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"sort"
)

var (
	ErrFreezingInfo = errors.New("FreezingInfo must be grater than 0")
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

func ParserListProductToDTO(p map[int]model.Product) []ProductDTO {
	list := make([]ProductDTO, 0, len(p))
	for _, v := range p {
		list = append(list, ParserProductToDTO(v))
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}

func ParserProductToDTO(p model.Product) ProductDTO {
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

func ValidAndParserDTO(dto ProductDTO, productToParser *model.Product) error {

	if dto.ProductCode == "" {
		return errors.New("Invalid ProductCode")
	}
	if dto.Description == "" {
		return errors.New("description Can't be empty")
	}

	productTypeId, errID := domain.NewId(dto.ProductTypeID)
	if errID != nil {
		return errID
	}

	sellerId, errID := domain.NewId(dto.SellerID)
	if errID != nil {
		return errID
	}

	if dto.ProductCode == "" {
		return errors.New("Invalid ProductCode")
	}

	if errDimensions := validDimensions(dto.Height, dto.Length, dto.NetWeight, dto.Width); errDimensions != nil {
		return errDimensions
	}

	if errFreezingInfo := validFreezingInfo(dto.ExpirationRate, dto.FreezingRate); errFreezingInfo != nil {
		return errFreezingInfo
	}

	*productToParser = parseDTOProduct(productTypeId, sellerId, dto)
	return nil
}

func validDimensions(Height, Length, NetWeight, Width float64) error {
	if Height <= 0 {
		return errors.New("Invalid Height")
	}
	if Length <= 0 {
		return errors.New("Invalid Length")
	}
	if NetWeight <= 0 {
		return errors.New("Invalid NetWeight")
	}
	if Width <= 0 {
		return errors.New("Invalid Width")
	}
	return nil
}

func validFreezingInfo(ExpirationRate, FreezingRate float64) error {
	if ExpirationRate <= 0 {
		return ErrFreezingInfo
	}
	if FreezingRate <= 0 {
		return ErrFreezingInfo
	}

	//TODO agregar validaciones para RecommendedFreezingTemperature por rangos de temp
	return nil
}

func parseDTOProduct(productTypeID, sellerID domain.Id, product ProductDTO) model.Product {
	id, _ := domain.NewId(product.ID)
	return model.Product{
		ID:            id,
		Description:   product.Description,
		ProductCode:   product.ProductCode,
		ProductTypeID: productTypeID,
		SellerID:      sellerID,
		FreezingInfo: model.FreezingInfo{
			ExpirationRate:                 product.ExpirationRate,
			FreezingRate:                   product.FreezingRate,
			RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		},
		Dimensions: model.Dimensions{
			Height:    product.Height,
			Length:    product.Length,
			NetWeight: product.NetWeight,
			Width:     product.Width,
		},
	}

}

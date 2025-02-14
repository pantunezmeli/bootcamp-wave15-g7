package product

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

var (
	ErrFreezingInfo = errors.New("FreezingInfo must be greater than 0")
)

func ValidAndParserDTO(dto ProductDTO, productToParser *models.Product) error {

	if dto.ProductCode == "" {
		return errors.New("invalid ProductCode")
	}
	if dto.Description == "" {
		return errors.New("description Can't be empty")
	}

	productTypeId, errID := value_objects.NewProductTypeId(dto.ProductTypeID)
	if errID != nil {
		return errID
	}

	sellerId, errID := value_objects.NewSellerId(dto.SellerID)
	if errID != nil {
		return errID
	}

	if dto.ProductCode == "" {
		return errors.New("invalid ProductCode")
	}

	if errDimensions := ValidDimensions(dto.Height, dto.Length, dto.NetWeight, dto.Width); errDimensions != nil {
		return errDimensions
	}

	if errFreezingInfo := ValidFreezingInfo(dto.ExpirationRate, dto.FreezingRate); errFreezingInfo != nil {
		return errFreezingInfo
	}

	*productToParser = ParseDTOProduct(productTypeId, sellerId, dto)
	return nil
}

func ValidDimensions(Height, Length, NetWeight, Width float64) error {
	if Height <= 0 {
		return errors.New("invalid Height")
	}
	if Length <= 0 {
		return errors.New("invalid Length")
	}
	if NetWeight <= 0 {
		return errors.New("invalid NetWeight")
	}
	if Width <= 0 {
		return errors.New("invalid Width")
	}
	return nil
}

func ValidFreezingInfo(ExpirationRate, FreezingRate float64) error {
	if ExpirationRate <= 0 {
		return ErrFreezingInfo
	}
	if FreezingRate <= 0 {
		return ErrFreezingInfo
	}

	//TODO add valid RecommendedFreezingTemperature by range
	return nil
}

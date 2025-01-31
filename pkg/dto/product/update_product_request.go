package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"reflect"
)

type UpdateProductRequest struct {
	Description                    *string  `json:"description"`
	ExpirationRate                 *float64 `json:"expiration_rate"`
	FreezingRate                   *float64 `json:"freezing_rate"`
	Height                         *float64 `json:"height"`
	Length                         *float64 `json:"length"`
	NetWeight                      *float64 `json:"netweight"`
	ProductCode                    *string  `json:"product_code"`
	RecommendedFreezingTemperature *float64 `json:"recommended_freezing_temperature"`
	Width                          *float64 `json:"width"`
	ProductTypeID                  *int     `json:"product_type_id"`
	SellerID                       *int     `json:"seller_id"`
}

func PatchProduct(patch UpdateProductRequest, productEntity *models.Product) error {
	dtoToPatch := ParserProductToDTO(*productEntity)

	patchVal := reflect.ValueOf(patch)
	dtoVal := reflect.ValueOf(&dtoToPatch).Elem()

	for i := 0; i < patchVal.NumField(); i++ {
		fieldName := patchVal.Type().Field(i).Name
		patchField := patchVal.Field(i)

		if !patchField.IsNil() {
			dtoField := dtoVal.FieldByName(fieldName)
			if dtoField.IsValid() && dtoField.CanSet() {
				dtoField.Set(patchField.Elem())
			}
		}
	}

	return ValidAndParserDTO(dtoToPatch, productEntity)
}

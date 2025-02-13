package locality

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	locality_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/locality"
	province_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/province"
)

type LocalityRequest struct {
	Name *string `json:"locality_name"`
	ProvinceId *int `json:"province_id"`
}

type LocalityDoc struct {
	Id locality_vo.LocalityId
	Name locality_vo.LocalityName
	ProvinceId province_vo.ProvinceId
}

type SellerReport struct {
	LocalityId int `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount int `json:"sellers_count"`
}

func ParseRequestToModel(localityRequest LocalityRequest) (localityModel models.Locality, err error){
	name, err := locality_vo.NewLocalityName(*localityRequest.Name)
	if err != nil {
		return
	}

	provinceId, err := province_vo.NewProvinceId(*localityRequest.ProvinceId)
	if err != nil {
		return
	}

	localityModel = models.Locality{
		LocalityAttributes: models.LocalityAttributes{
			Name: name,
		},
		ProvinceId: provinceId,
	}

	return
}

func ParseModelToResponse(localityModel models.Locality) (localityDto LocalityDoc){
	localityDto = LocalityDoc{
			Id: localityModel.Id,
			Name: localityModel.Name,
			ProvinceId: localityModel.ProvinceId,
		}
	return
}

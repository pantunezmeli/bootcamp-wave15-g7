package models

import (
	province_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/province"
	country_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/country"

)

type Province struct {
	Id province_vo.ProvinceId
	ProvinceAttributes
	CountryId country_vo.CountryId
}

type ProvinceAttributes struct {
	Name province_vo.ProvinceName
}
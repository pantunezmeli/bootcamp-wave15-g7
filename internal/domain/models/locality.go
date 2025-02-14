package models

import (
	locality_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/locality"
	province_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/province"
)

type Locality struct {
	Id locality_vo.LocalityId
	LocalityAttributes
	ProvinceId province_vo.ProvinceId
}

type LocalityAttributes struct {
	Name locality_vo.LocalityName
}
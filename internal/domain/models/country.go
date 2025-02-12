package models

import (
	country_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/country"
)

type Country struct {
	Id country_vo.CountryId
	CountryAttributes
}

type CountryAttributes struct {
	CountryName country_vo.CountryName
}
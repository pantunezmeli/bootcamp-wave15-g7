package models

import (
	d "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
)

type WareHouse struct {
	Id                 d.Id
	WareHouseCode      d.WareHouseCode
	Address            d.Address
	Telephone          d.Telephone
	MinimunCapacity    d.MinimunCapacity
	MinimunTemperature d.MinimunTemperature
}

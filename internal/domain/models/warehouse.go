package models

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type WareHouse struct {
	Id                value_objects.Id
	WareHouseCode     value_objects.WareHouseCode
	Address           value_objects.Address
	Telephone         value_objects.Telephone
	MinimunCapacity   value_objects.MinimunCapacity
	MinimunTemperature value_objects.MinimunTemperature
}

package models

import (
	vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type WareHouse struct {
	Id            vo.IDWarehouse
	WareHouseCode vo.WareHouseCode
	Address       vo.WarehouseAddress
	Telephone     vo.WarehouseTelephone
	LocalityId    vo.LocalityID
}

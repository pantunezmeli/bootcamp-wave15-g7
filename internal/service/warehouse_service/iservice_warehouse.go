package warehouse_service

import (
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type IWareHouseService interface {
	// ! 1)
	FindAll() (w map[int]dto.WareHouseDoc, err error)

	// ! 2)
	GetWareHouseById(id int) (w dto.WareHouseDoc, err error)

	// ! 3)
	AddWareHouse(req dto.WareHouseDoc) (err error)

	// ! 4)
	EditWareHouse(id int, req dto.WareHouseDoc) (wh dto.WareHouseDoc, err error)
}

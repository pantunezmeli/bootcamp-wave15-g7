package warehouse_service

import (
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type IWareHouseService interface {
	// ! 1)
	FindAll() (v map[int]dto.WareHouseDoc, err error)
}

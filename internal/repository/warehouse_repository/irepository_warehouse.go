package warehouse

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IWareHouseRepository interface {
	GetAllWareHouses() (w map[int]models.WareHouse, err error)
}

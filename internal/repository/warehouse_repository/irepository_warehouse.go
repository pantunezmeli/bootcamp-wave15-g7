package warehouse

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IWareHouseRepository interface {

	//! 1)
	GetAllWareHouses() (w map[int]models.WareHouse, err error)

	//! 2)
	GetWareHouseById(id int) (wh models.WareHouse, exists bool)

	//! 3)
	GetWareHouseByCode(code string) (wh models.WareHouse, exists bool)
	CreateNewWareHouse(wh models.WareHouse) (err error)

	//! 4)
	UpdateWarehouse(wh models.WareHouse) (err error)
}

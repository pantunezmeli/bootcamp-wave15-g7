package warehouse

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IWareHouseRepository interface {

	//! 1)
	GetAllWareHouses() (w []models.WareHouse, err error)

	//! 2)
	GetWareHouseById(id int) (wh models.WareHouse, err error)

	//! 3)
	GetWareHouseByCode(code string) (wh models.WareHouse, err error)
	CreateNewWareHouse(wh models.WareHouse) (w models.WareHouse, err error)

	// //! 4)
	UpdateWarehouse(wh models.WareHouse) (err error)

	//! 5)
	DeleteWarehouse(id int) (err error)
}

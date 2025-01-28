package warehouse

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	loader "github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader/warehouse_loader"
)

type WareHouseRepository struct {
	db     map[int]models.WareHouse
	loader loader.IWareHouseLoader
}

// Crea una nueva 'base de datos'
func NewWareHouseRepository(db map[int]models.WareHouse, loader loader.IWareHouseLoader) *WareHouseRepository {
	return &WareHouseRepository{
		db:     db,
		loader: loader,
	}
}

// ! 1)
func (r *WareHouseRepository) GetAllWareHouses() (w map[int]models.WareHouse, err error) {
	w = make(map[int]models.WareHouse)
	for key, value := range r.db {
		w[key] = value
	}
	return
}

// ! 2)
func (r *WareHouseRepository) GetWareHouseById(id int) (wh models.WareHouse, exists bool) {
	wh, exists = r.db[id]
	return wh, exists
}

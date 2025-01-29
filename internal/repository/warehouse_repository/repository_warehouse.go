package warehouse

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	loader "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/warehouse_storage"
)

type WareHouseRepository struct {
	db     map[int]models.WareHouse
	loader loader.IWareHouseStorage
}

// Crea una nueva 'base de datos'
func NewWareHouseRepository(db map[int]models.WareHouse, loader loader.IWareHouseStorage) *WareHouseRepository {
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

// ! 3)
func (r *WareHouseRepository) GetWareHouseByCode(code string) (wh models.WareHouse, exists bool) {
	for _, wh := range r.db {
		if wh.WareHouseCode.GetWareHouseCode() == code {
			return wh, true
		}
	}
	return wh, false
}

func (r *WareHouseRepository) CreateNewWareHouse(wh models.WareHouse) (err error) {
	r.db[wh.Id.GetId()] = wh
	return nil
}

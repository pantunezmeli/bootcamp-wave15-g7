package warehouse

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	storage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/warehouse_storage"
)

var (
	ErrLoadingData           = errors.New("error loading data")
	ErrWareHouseNotFound     = errors.New("warehouse not found")
	ErrWareHouseCodeNotFound = errors.New("warehouse code not found")
)

type WareHouseRepository struct {
	//db     map[int]models.WareHouse
	storage storage.IWareHouseStorage
}

// Crea una nueva 'base de datos'
func NewWareHouseRepository(storage storage.IWareHouseStorage) *WareHouseRepository {
	return &WareHouseRepository{
		storage: storage,
	}
}

// ! 1)
func (r *WareHouseRepository) GetAllWareHouses() (map[int]models.WareHouse, error) {
	// Load data
	db, err := r.storage.Load()

	// Fail load
	if err != nil {
		return map[int]models.WareHouse{}, err
	}

	return db, nil
}

// ! 2)
func (r *WareHouseRepository) GetWareHouseById(id int) (wh models.WareHouse, err error) {
	db, err := r.storage.Load()
	if err != nil {
		return models.WareHouse{}, err
	}
	wh, exists := db[id]
	if !exists {
		return models.WareHouse{}, ErrWareHouseNotFound
	}
	return
}

// ! 3)
func (r *WareHouseRepository) GetWareHouseByCode(code string) (wh models.WareHouse, err error) {
	// Load data
	db, err := r.storage.Load()

	// Fail load
	if err != nil {
		return models.WareHouse{}, ErrLoadingData
	}

	// Check if exists Warehouse_code
	for _, wh := range db {
		if wh.WareHouseCode.GetWareHouseCode() == code {
			return wh, nil
		}
	}

	// Return empty
	return models.WareHouse{}, ErrWareHouseCodeNotFound
}

func (r *WareHouseRepository) CreateNewWareHouse(wh models.WareHouse) (err error) {
	// Load data
	db, err := r.storage.Load()

	// Fail load
	if err != nil {
		return ErrLoadingData
	}

	// Add wh to memory
	db[wh.Id.GetId()] = wh

	// Save data
	r.storage.Save(db)

	return nil
}

// ! 4)
func (r *WareHouseRepository) UpdateWarehouse(wh models.WareHouse) (err error) {

	// Load data
	db, err := r.storage.Load()
	if err != nil {
		return ErrLoadingData
	}

	// Save on memory
	db[wh.Id.GetId()] = wh

	// Save on file
	r.storage.Save(db)
	return nil
}

// ! 5)
func (r *WareHouseRepository) DeleteWarehouse(id int) (err error) {

	// Load data
	db, err := r.storage.Load()
	if err != nil {
		return ErrLoadingData
	}

	// Delete on memory
	delete(db, id)

	// Save on file
	r.storage.Save(db)

	return nil
}

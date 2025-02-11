package warehouse

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

var (
	ErrLoadingData           = errors.New("error loading data")
	ErrWareHouseNotFound     = errors.New("warehouse not found")
	ErrWareHouseCodeNotFound = errors.New("warehouse code not found")
)

type WareHouseRepository struct {
	db *sql.DB // Contiene una base de datos
}

// Inyección de la base de datos
func NewWareHouseRepository(db *sql.DB) *WareHouseRepository {
	return &WareHouseRepository{
		db: db,
	}
}

// ! 1)
func (r *WareHouseRepository) GetAllWareHouses() ([]models.WareHouse, error) {
	query := "SELECT id, warehouse_code, address, telephone, minimunCapacity, minimunTemperature FROM warehouse"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []models.WareHouse

	// Iteración sobre los resultados
	for rows.Next() {
		var warehouse models.WareHouse

		err := rows.Scan(
			&warehouse.Id,
			&warehouse.WareHouseCode,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.MinimunCapacity,
			&warehouse.MinimunTemperature,
		)
		if err != nil {
			return nil, err
		}

		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

// ! 2)
func (r *WareHouseRepository) GetWareHouseById(id int) (models.WareHouse, error) {

	var warehouse models.WareHouse
	query := "SELECT id, wareHouseCode, address, telephone, minimunCapacity, minimunTemperature FROM warehouse WHERE id = ?"
	row := r.db.QueryRow(query, id)

	err := row.Scan(
		&warehouse.Id,
		&warehouse.WareHouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.MinimunCapacity,
		&warehouse.MinimunTemperature,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.WareHouse{}, fmt.Errorf("warehouse with id %d not found", id)
		}
		return models.WareHouse{}, err
	}
	return warehouse, nil
}

// ! 3)
func (r *WareHouseRepository) GetWareHouseByCode(code string) (models.WareHouse, error) {
	var warehouse models.WareHouse
	query := "SELECT * FROM warehouse WHERE warehouse.wareHouseCode = ?"
	row := r.db.QueryRow(query, code)

	err := row.Scan(
		&warehouse.Id,
		&warehouse.WareHouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.MinimunCapacity,
		&warehouse.MinimunTemperature,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.WareHouse{}, ErrWareHouseCodeNotFound
		}
		return models.WareHouse{}, err
	}
	return warehouse, nil
}

func (r *WareHouseRepository) CreateNewWareHouse(warehouse models.WareHouse) (models.WareHouse, error) {

	query := "INSERT INTO warehouse (wareHouseCode, address, telephone, minimunCapacity, minimunTemperature) VALUES (?, ?, ?, ?, ?)"

	result, err := r.db.Exec(query,
		warehouse.WareHouseCode.GetWareHouseCode(),
		warehouse.Address.GetAddress(),
		warehouse.Telephone.GetTelephone(),
		warehouse.MinimunCapacity.GetMinimunCapacity(),
		warehouse.MinimunTemperature.GetMinimunTemperature(),
	)

	if err != nil {
		return models.WareHouse{}, err // Err
	}

	// Obtener el último ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return models.WareHouse{}, err // ErrGettingLastInsertedID
	}

	// Conversión del ID a objeto Value_object
	newIdObj, err := value_objects.NewId(int(lastInsertID))
	if err != nil {
		return models.WareHouse{}, err // ErrConvertingIDToValueObject
	}

	warehouse.Id = newIdObj

	return warehouse, nil
}

// ! 4)
func (r *WareHouseRepository) UpdateWarehouse(warehouse models.WareHouse) error {

	query := "UPDATE warehouse SET warehouse_code = ?, address = ?, telephone = ?, minimunCapacity = ?, minimunTemperature = ? WHERE id = ?"

	_, err := r.db.Exec(query,
		warehouse.WareHouseCode.GetWareHouseCode(),
		warehouse.Address.GetAddress(),
		warehouse.Telephone.GetTelephone(),
		warehouse.MinimunCapacity.GetMinimunCapacity(),
		warehouse.MinimunTemperature.GetMinimunTemperature(),
		warehouse.Id.GetId(),
	)
	if err != nil {
		return err // ErrExecutingDB
	}

	return nil
}

// ! 5)
func (r *WareHouseRepository) DeleteWarehouse(id int) (err error) {
	query := "DELETE FROM warehouse WHERE id = ?"

	_, err = r.db.Exec(query,
		id)
	if err != nil {
		return err
	}

	return nil
}

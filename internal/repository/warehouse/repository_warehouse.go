package warehouse

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	customErrors "github.com/pantunezmeli/bootcamp-wave15-g7/internal/errors"
)

var (
	ErrLoadingData           = errors.New("error loading data")
	ErrWareHouseNotFound     = errors.New("warehouse not found")
	ErrWareHouseCodeNotFound = errors.New("warehouse code not found")
)

type WareHouseRepository struct {
	db *sql.DB
}

func NewWareHouseRepository(db *sql.DB) *WareHouseRepository {
	return &WareHouseRepository{
		db: db,
	}
}

// ! 1)
func (r *WareHouseRepository) GetAllWareHouses() ([]models.WareHouse, error) {
	query := "SELECT id, warehouse_code, address, telephone, locality_id FROM warehouses"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, customErrors.ErrDBGenericError
	}
	defer rows.Close()

	var warehouses []models.WareHouse

	for rows.Next() {
		var warehouse models.WareHouse

		err := rows.Scan(
			&warehouse.Id,
			&warehouse.WareHouseCode,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.LocalityId,
		)
		if err != nil {
			return nil, customErrors.ErrMappingData
		}

		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

// ! 2)
func (r *WareHouseRepository) GetWareHouseById(id int) (models.WareHouse, error) {

	var warehouse models.WareHouse
	query := "SELECT id, warehouse_code, address, telephone, locality_id FROM warehouses WHERE id = ?"
	row := r.db.QueryRow(query, id)

	err := row.Scan(
		&warehouse.Id,
		&warehouse.WareHouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.LocalityId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.WareHouse{}, customErrors.ErrWarehouseNotFound
		}
		return models.WareHouse{}, customErrors.ErrMappingData
	}
	return warehouse, nil
}

// ! 3)

func (r *WareHouseRepository) CreateNewWareHouse(warehouse models.WareHouse) (models.WareHouse, error) {

	query := "INSERT INTO warehouses (warehouse_code, address, telephone, locality_id) VALUES (?, ?, ?, ?)"

	result, err := r.db.Exec(query,
		warehouse.WareHouseCode,
		warehouse.Address,
		warehouse.Telephone,
		warehouse.LocalityId,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1452:
				fmt.Println("Entré en el caso de que no existe la foranea")
				return models.WareHouse{}, customErrors.ErrForeignKeyViolation
			case 1062:
				fmt.Println("Entré en el ya existe un warehouse code igual")
				return models.WareHouse{}, customErrors.ErrWarehouseCodeDuplicate
			default:
				fmt.Println("Entré en el caso de que tengo un error genérico de BD")
				return models.WareHouse{}, customErrors.ErrDBGenericError
			}
		}
		fmt.Println("Error insertando el elemento en la BD")
		return models.WareHouse{}, customErrors.ErrInsertingData
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return models.WareHouse{}, customErrors.ErrGettingLastID
	}

	newIdObj, err := value_objects.NewWarehouseId(int(lastInsertID))
	if err != nil {
		return models.WareHouse{}, customErrors.ErrConvertingID
	}

	warehouse.Id = newIdObj

	return warehouse, nil
}

// ! 4)
func (r *WareHouseRepository) UpdateWarehouse(warehouse models.WareHouse) error {

	query := "UPDATE warehouses SET warehouse_code = ?, address = ?, telephone = ?, locality_id = ? WHERE id = ?"

	_, err := r.db.Exec(query,
		warehouse.WareHouseCode,
		warehouse.Address,
		warehouse.Telephone,
		warehouse.Id,
		warehouse.LocalityId,
	)
	if err != nil {
		return err // ErrExecutingDB
	}

	return nil
}

// ! 5)
func (r *WareHouseRepository) DeleteWarehouse(id int) (err error) {
	query := "DELETE FROM warehouses WHERE id = ?"

	_, err = r.db.Exec(query,
		id)
	if err != nil {
		return err
	}

	return nil
}

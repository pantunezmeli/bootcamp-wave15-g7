package warehouse

import (
	"errors"
	"fmt"

	customErrors "github.com/pantunezmeli/bootcamp-wave15-g7/internal/errors"
	repository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/warehouse"
)

var (
	ErrWareHouseNotFound          = errors.New("warehouse not found")
	ErrWareHouseCodeAlreadyExists = errors.New("warehouse with that code already exists")
	ErrInvalidData                = errors.New("invalid warehouse data")
	ErrInvalidIdGenerated         = errors.New("invalid id generated")
)

type WarehouseService struct {
	rp repository.IWareHouseRepository
}

func NewWareHouseService(rp repository.IWareHouseRepository) *WarehouseService {
	return &WarehouseService{rp: rp}
}

// ! 1)
func (s *WarehouseService) FindAll() (w map[int]dto.WareHouseDoc, err error) {

	warehouses, err := s.rp.GetAllWareHouses()

	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrMappingData):
			return map[int]dto.WareHouseDoc{}, customErrors.ErrConvertion{Err: err}
		default:
			return map[int]dto.WareHouseDoc{}, customErrors.ErrDatabase{Err: fmt.Errorf("unexpected error: %w", err)}
		}
	}
	w = dto.MapWarehousesToDTO(warehouses)

	return
}

// ! 2)
func (s *WarehouseService) GetWareHouseById(id int) (w dto.WareHouseDoc, err error) {
	wh, err := s.rp.GetWareHouseById(id)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrMappingData):
			return dto.WareHouseDoc{}, customErrors.ErrConvertion{Err: err}
		case errors.Is(err, customErrors.ErrWarehouseNotFound):
			return dto.WareHouseDoc{}, customErrors.ErrNotFound{Err: err}
		default:
			return dto.WareHouseDoc{}, customErrors.ErrDatabase{Err: fmt.Errorf("unexpected error: %w", err)}
		}
	}

	whDTO, err := dto.WareHouseDoc{}.ConvertToDTO(wh)
	if err != nil {
		return dto.WareHouseDoc{}, customErrors.ErrConvertion{Err: err}
	}
	return whDTO, nil
}

// ! 3)
func (s *WarehouseService) AddWareHouse(req dto.WareHouseDoc) (dto.WareHouseDoc, error) {

	warehouse, err := req.ConvertToModel(req)
	if err != nil {
		return dto.WareHouseDoc{}, ErrInvalidParameter{Parameter: err.Error()}
	}

	createdWarehouse, err := s.rp.CreateNewWareHouse(warehouse)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrForeignKeyViolation):
			return dto.WareHouseDoc{}, customErrors.ErrForeignKey{Err: err}
		case errors.Is(err, customErrors.ErrWarehouseCodeDuplicate):
			return dto.WareHouseDoc{}, customErrors.ErrDuplicate{Err: err}
		case errors.Is(err, customErrors.ErrDBGenericError):
			return dto.WareHouseDoc{}, customErrors.ErrDatabase{Err: err}
		case errors.Is(err, customErrors.ErrInsertingData):
			return dto.WareHouseDoc{}, customErrors.ErrDatabase{Err: err}
		case errors.Is(err, customErrors.ErrGettingLastID), errors.Is(err, customErrors.ErrConvertingID):
			return dto.WareHouseDoc{}, customErrors.ErrDatabase{Err: err}
		default:
			return dto.WareHouseDoc{}, customErrors.ErrDatabase{Err: fmt.Errorf("unexpected error: %w", err)}
		}
	}

	whDTO, err := dto.WareHouseDoc{}.ConvertToDTO(createdWarehouse)
	if err != nil {
		return dto.WareHouseDoc{}, err
	}

	return whDTO, nil
}

// ! 4)
func (s *WarehouseService) EditWareHouse(id int, req dto.WareHouseDoc) (whDTO dto.WareHouseDoc, err error) {

	existingWarehouse, err := s.rp.GetWareHouseById(id)
	if err != nil {
		return dto.WareHouseDoc{}, ErrWareHouseNotFound
	}

	warehouse, err := req.ConvertToModelPatch(req, existingWarehouse)
	if err != nil {
		return dto.WareHouseDoc{}, ErrInvalidParameter{Parameter: err.Error()}
	}

	// if warehouse.WareHouseCode != existingWarehouse.WareHouseCode {
	// 	_, err := s.rp.GetWareHouseByCode(string(warehouse.WareHouseCode))
	// 	if err == nil {
	// 		return dto.WareHouseDoc{}, ErrWareHouseCodeAlreadyExists
	// 	}
	// }

	warehouse.Id = existingWarehouse.Id

	err = s.rp.UpdateWarehouse(warehouse)
	if err != nil {
		return dto.WareHouseDoc{}, err
	}

	whDTO, err = dto.WareHouseDoc{}.ConvertToDTO(warehouse)
	if err != nil {
		return dto.WareHouseDoc{}, err
	}
	return
}

// ! 5)
func (s *WarehouseService) DeleteWarehouse(id int) error {

	_, err := s.rp.GetWareHouseById(id)
	if err != nil {
		return ErrWareHouseNotFound
	}

	err = s.rp.DeleteWarehouse(id)
	if err != nil {
		return err
	}

	return nil
}

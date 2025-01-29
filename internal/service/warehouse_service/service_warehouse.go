package warehouse_service

import (
	"errors"
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	repository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse_repository"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
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

	wareHouses, err := s.rp.GetAllWareHouses()
	if err != nil {
		return map[int]dto.WareHouseDoc{}, err
	}
	return s.MapWareHouseToDTO(wareHouses), nil
}

// ! 2)
func (s *WarehouseService) GetWareHouseById(id int) (w dto.WareHouseDoc, err error) {
	wh, exists := s.rp.GetWareHouseById(id)
	if !exists {
		return dto.WareHouseDoc{}, ErrWareHouseNotFound
	}

	// Convert to DTO
	whDTO, err := dto.WareHouseDoc{}.ConvertToDTO(wh)
	if err != nil {
		return dto.WareHouseDoc{}, err // CAPTURAR EL ERROR
	}
	return whDTO, nil
}

// ! 3)
func (s *WarehouseService) AddWareHouse(req dto.WareHouseDoc) (wh dto.WareHouseDoc, err error) {

	// Validation and convert to model
	warehouse, err := req.ConvertToModel(req)
	if err != nil {
		return dto.WareHouseDoc{}, fmt.Errorf("failed to convert warehouse: %w", err)
	}

	// Validation of warehouse code
	_, exists := s.rp.GetWareHouseByCode(warehouse.WareHouseCode.GetWareHouseCode())
	if exists {
		return dto.WareHouseDoc{}, ErrWareHouseCodeAlreadyExists
	}

	// Generation of new Id
	warehouses, err := s.rp.GetAllWareHouses()
	if err != nil {
		return dto.WareHouseDoc{}, err
	}

	var newId = 0
	for id := range warehouses {
		if id > newId {
			newId = id
		}
	}
	newId++

	// Asignation of Id
	newIdObj, err := domain.NewId(newId)
	if err != nil {
		return dto.WareHouseDoc{}, ErrInvalidIdGenerated
	}

	warehouse.Id = newIdObj

	// Call the repository
	err = s.rp.CreateNewWareHouse(warehouse)
	if err != nil {
		return dto.WareHouseDoc{}, err
	}

	// Convert model to DTO
	wh, err = wh.ConvertToDTO(warehouse)
	if err != nil {
		return dto.WareHouseDoc{}, err
	}

	return
}

func (s *WarehouseService) MapWareHouseToDTO(w map[int]models.WareHouse) (r map[int]dto.WareHouseDoc) {
	r = make(map[int]dto.WareHouseDoc)
	for id, wh := range w {
		r[id] = dto.WareHouseDoc{
			Id:                 wh.Id.GetId(),
			WareHouseCode:      wh.WareHouseCode.GetWareHouseCode(),
			Address:            wh.Address.GetAddress(),
			Telephone:          wh.Telephone.GetTelephone(),
			MinimunCapacity:    wh.MinimunCapacity.GetMinimunCapacity(),
			MinimunTemperature: wh.MinimunTemperature.GetMinimunTemperature(),
		}
	}
	return
}

// ! 4)
func (s *WarehouseService) EditWareHouse(id int, req dto.WareHouseDoc) (wh dto.WareHouseDoc, err error) {
	// Get previous instance
	existingWarehouse, exists := s.rp.GetWareHouseById(id)
	if !exists {
		return dto.WareHouseDoc{}, ErrWareHouseNotFound
	}

	// Validation and covert to model
	warehouse, err := req.ConvertToModelPatch(req, existingWarehouse)
	if err != nil {
		return dto.WareHouseDoc{}, fmt.Errorf("failed to convert warehouse: %w", err)
	}

	warehouse.Id = existingWarehouse.Id

	// Verify warehouse_code
	whCodeExisting, exists := s.rp.GetWareHouseByCode(warehouse.WareHouseCode.GetWareHouseCode())
	if exists && whCodeExisting.Id.GetId() != id {
		return dto.WareHouseDoc{}, ErrWareHouseCodeAlreadyExists
	}

	// Call Repository
	err = s.rp.UpdateWarehouse(warehouse)
	if err != nil {
		return dto.WareHouseDoc{}, err
	}

	// Convert to DTO
	wh, err = wh.ConvertToDTO(warehouse)
	if err != nil {
		return dto.WareHouseDoc{}, err
	}
	return
}

// ! 5)
func (s *WarehouseService) DeleteWarehouse(id int) error {
	// Get previous instance
	_, exists := s.rp.GetWareHouseById(id)
	if !exists {
		return ErrWareHouseNotFound
	}

	// Call repository
	err := s.rp.DeleteWarehouse(id)
	if err != nil {
		return err
	}
	return nil
}

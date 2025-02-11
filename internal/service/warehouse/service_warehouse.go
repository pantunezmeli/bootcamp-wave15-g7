package warehouse

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
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

	wareHouses, err := s.rp.GetAllWareHouses()

	if err != nil {
		return map[int]dto.WareHouseDoc{}, err
	}
	w = s.MapWareHouseToDTO(wareHouses)

	return
}

// TODO mover al DTO
func (s *WarehouseService) MapWareHouseToDTO(w []models.WareHouse) (r map[int]dto.WareHouseDoc) {
	r = make(map[int]dto.WareHouseDoc)
	for _, wh := range w {
		r[wh.Id.GetId()] = dto.WareHouseDoc{
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

// ! 2)
func (s *WarehouseService) GetWareHouseById(id int) (w dto.WareHouseDoc, err error) {
	wh, err := s.rp.GetWareHouseById(id)
	if err != nil {
		return dto.WareHouseDoc{}, ErrWareHouseNotFound
	}

	whDTO, err := dto.WareHouseDoc{}.ConvertToDTO(wh)
	if err != nil {
		return dto.WareHouseDoc{}, err
	}
	return whDTO, nil
}

// ! 3)
func (s *WarehouseService) AddWareHouse(req dto.WareHouseDoc) (dto.WareHouseDoc, error) {

	warehouse, err := req.ConvertToModel(req)
	if err != nil {
		return dto.WareHouseDoc{}, ErrInvalidParameter{Parameter: err.Error()}
	}

	_, err = s.rp.GetWareHouseByCode(warehouse.WareHouseCode.GetWareHouseCode())
	if err == nil {
		return dto.WareHouseDoc{}, ErrWareHouseCodeAlreadyExists
	}

	if err != repository.ErrWareHouseCodeNotFound {
		return dto.WareHouseDoc{}, err
	}

	createdWarehouse, err := s.rp.CreateNewWareHouse(warehouse)
	if err != nil {
		return dto.WareHouseDoc{}, err
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

	if warehouse.WareHouseCode.GetWareHouseCode() != existingWarehouse.WareHouseCode.GetWareHouseCode() {
		_, err := s.rp.GetWareHouseByCode(warehouse.WareHouseCode.GetWareHouseCode())
		if err == nil {
			return dto.WareHouseDoc{}, ErrWareHouseCodeAlreadyExists
		}
	}

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

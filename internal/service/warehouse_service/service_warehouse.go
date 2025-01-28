package warehouse_service

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	repository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse_repository"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

var (
	ErrWareHouseNotFound = errors.New("warehouse not found")
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

	return dto.WareHouseDoc{
		Id:                 wh.Id.GetId(),
		WareHouseCode:      wh.WareHouseCode.GetWareHouseCode(),
		Address:            wh.Address.GetAddress(),
		Telephone:          wh.Telephone.GetTelephone(),
		MinimunCapacity:    wh.MinimunCapacity.GetMinimunCapacity(),
		MinimunTemperature: wh.MinimunTemperature.GetMinimunTemperature(),
	}, nil
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

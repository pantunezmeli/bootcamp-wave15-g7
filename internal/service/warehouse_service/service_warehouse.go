package warehouse_service

import (
	repository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse_repository"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
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

	// Mapped
	result := make(map[int]dto.WareHouseDoc)
	for id, wh := range wareHouses {
		result[id] = dto.WareHouseDoc{
			Id:                 wh.Id.GetId(),
			WareHouseCode:      wh.WareHouseCode.GetWareHouseCode(),
			Address:            wh.Address.GetAddress(),
			Telephone:          wh.Telephone.GetTelephone(),
			MinimunCapacity:    wh.MinimunCapacity.GetMinimunCapacity(),
			MinimunTemperature: wh.MinimunTemperature.GetMinimunTemperature(),
		}
	}

	return result, nil
}

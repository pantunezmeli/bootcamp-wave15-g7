package warehouse

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type WareHouseDoc struct {
	Id            int    `json:"id"`
	WareHouseCode string `json:"warehouse_code"`
	Address       string `json:"address"`
	Telephone     string `json:"telephone"`
	LocalityId    int    `json:"locality_id"`
}

func (w WareHouseDoc) ConvertToModel(req WareHouseDoc) (models.WareHouse, error) {

	wareHouseCode, err := value_objects.NewWareHouseCode(req.WareHouseCode)
	if err != nil {
		return models.WareHouse{}, err
	}

	address, err := value_objects.NewWarehouseAddress(req.Address)
	if err != nil {
		return models.WareHouse{}, err
	}

	telephone, err := value_objects.NewWarehouseTelephone(req.Telephone)
	if err != nil {
		return models.WareHouse{}, err
	}

	localityId, err := value_objects.NewWarehouseLocalityID(req.LocalityId)
	if err != nil {
		return models.WareHouse{}, err
	}

	wh := models.WareHouse{
		WareHouseCode: wareHouseCode,
		Address:       address,
		Telephone:     telephone,
		LocalityId:    localityId,
	}
	return wh, nil
}

func (w WareHouseDoc) ConvertToDTO(req models.WareHouse) (wh WareHouseDoc, err error) {
	return WareHouseDoc{
		Id:            int(req.Id),
		WareHouseCode: string(req.WareHouseCode),
		Address:       string(req.Address),
		Telephone:     string(req.Telephone),
		LocalityId:    int(req.LocalityId),
	}, nil
}

func (w WareHouseDoc) ConvertToModelPatch(req WareHouseDoc, existingWarehouse models.WareHouse) (models.WareHouse, error) {
	if req.WareHouseCode != "" {
		newCode, err := value_objects.NewWareHouseCode(req.WareHouseCode)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse code '%s': %w", req.WareHouseCode, err)
		}
		existingWarehouse.WareHouseCode = newCode
	}

	if req.Address != "" {
		newAddress, err := value_objects.NewWarehouseAddress(req.Address)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse address '%s': %w", req.Address, err)
		}
		existingWarehouse.Address = newAddress
	}
	if req.Telephone != "" {
		newTelephone, err := value_objects.NewWarehouseTelephone(req.Telephone)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse telephone '%s': %w", req.Telephone, err)
		}
		existingWarehouse.Telephone = newTelephone
	}

	if req.LocalityId > 0 {
		newLocalityId, err := value_objects.NewWarehouseLocalityID(req.LocalityId)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid locality id '%d': %w", req.LocalityId, err)
		}
		existingWarehouse.LocalityId = newLocalityId
	}

	return existingWarehouse, nil
}

func MapWarehousesToDTO(req []models.WareHouse) map[int]WareHouseDoc {
	r := make(map[int]WareHouseDoc)
	for _, wh := range req {
		r[int(wh.Id)] = WareHouseDoc{
			Id:            int(wh.Id),
			WareHouseCode: string(wh.WareHouseCode),
			Address:       string(wh.Address),
			Telephone:     string(wh.Telephone),
			LocalityId:    int(wh.LocalityId),
		}
	}
	return r
}

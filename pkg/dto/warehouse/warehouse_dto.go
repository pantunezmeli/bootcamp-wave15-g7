package warehouse

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type WareHouseDoc struct {
	Id                 int    `json:"id"`
	WareHouseCode      string `json:"warehouse_code"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

func (w WareHouseDoc) ConvertToModel(req WareHouseDoc) (models.WareHouse, error) {

	wareHouseCode, err := value_objects.NewWareHouseCode(req.WareHouseCode)
	if err != nil {
		return models.WareHouse{}, err
	}

	address, err := value_objects.NewAddress(req.Address)
	if err != nil {
		return models.WareHouse{}, err
	}

	telephone, err := value_objects.NewTelephone(req.Telephone)
	if err != nil {
		return models.WareHouse{}, err
	}

	minCapacity, err := value_objects.NewMinimunCapacity(req.MinimunCapacity)
	if err != nil {
		return models.WareHouse{}, err
	}

	minTemp, err := value_objects.NewMinimunTemperature(req.MinimunTemperature)
	if err != nil {
		return models.WareHouse{}, err
	}

	wh := models.WareHouse{
		WareHouseCode:      wareHouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimunCapacity:    minCapacity,
		MinimunTemperature: minTemp,
	}
	return wh, nil
}

func (w WareHouseDoc) ConvertToDTO(req models.WareHouse) (wh WareHouseDoc, err error) {
	return WareHouseDoc{
		Id:                 value_objects.Id.GetId(req.Id),
		WareHouseCode:      value_objects.WareHouseCode.GetWareHouseCode(req.WareHouseCode),
		Address:            value_objects.Address.GetAddress(req.Address),
		Telephone:          value_objects.Telephone.GetTelephone(req.Telephone),
		MinimunCapacity:    value_objects.MinimunCapacity.GetMinimunCapacity(req.MinimunCapacity),
		MinimunTemperature: value_objects.MinimunTemperature.GetMinimunTemperature(req.MinimunTemperature),
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
		newAddress, err := value_objects.NewAddress(req.Address)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse address '%s': %w", req.Address, err)
		}
		existingWarehouse.Address = newAddress
	}
	if req.Telephone != "" {
		newTelephone, err := value_objects.NewTelephone(req.Telephone)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse telephone '%s': %w", req.Telephone, err)
		}
		existingWarehouse.Telephone = newTelephone
	}
	if req.MinimunCapacity > 0 {
		newMinCapacity, err := value_objects.NewMinimunCapacity(req.MinimunCapacity)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun capacity '%d': %w", req.MinimunCapacity, err)
		}
		existingWarehouse.MinimunCapacity = newMinCapacity
	}

	if req.MinimunTemperature > -100 {
		newMinTemp, err := value_objects.NewMinimunTemperature(req.MinimunTemperature)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun temperature '%d': %w", req.MinimunTemperature, err)
		}
		existingWarehouse.MinimunTemperature = newMinTemp
	}
	return existingWarehouse, nil
}

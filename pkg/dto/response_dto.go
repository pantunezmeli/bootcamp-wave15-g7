package dto

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
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

	wareHouseCode, err := domain.NewWareHouseCode(req.WareHouseCode)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse code '%s': %w", req.WareHouseCode, err)
	}

	address, err := domain.NewAddress(req.Address)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse address '%s': %w", req.Address, err)
	}

	telephone, err := domain.NewTelephone(req.Telephone)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse telephone '%s': %w", req.Telephone, err)
	}

	minCapacity, err := domain.NewMinimunCapacity(req.MinimunCapacity)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun capacity '%d': %w", req.MinimunCapacity, err)
	}

	minTemp, err := domain.NewMinimunTemperature(req.MinimunTemperature)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun temperature '%d': %w", req.MinimunTemperature, err)
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
		Id:                 domain.Id.GetId(req.Id),                                                 // Convertir value object Id a int
		WareHouseCode:      domain.WareHouseCode.GetWareHouseCode(req.WareHouseCode),                // Convertir value object WareHouseCode a string
		Address:            domain.Address.GetAddress(req.Address),                                  // Convertir value object Address a string
		Telephone:          domain.Telephone.GetTelephone(req.Telephone),                            // Convertir value object Telephone a string
		MinimunCapacity:    domain.MinimunCapacity.GetMinimunCapacity(req.MinimunCapacity),          // Convertir value object MinimunCapacity a int
		MinimunTemperature: domain.MinimunTemperature.GetMinimunTemperature(req.MinimunTemperature), // Convertir value object MinimunTemperature a int
	}, nil
}

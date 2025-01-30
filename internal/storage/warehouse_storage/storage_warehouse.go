package warehouse_storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type WareHouseJSONFile struct {
	path string
}

// Constructor
func NewWareHouseJSONFile(path string) *WareHouseJSONFile {
	return &WareHouseJSONFile{
		path: path,
	}
}

func (s *WareHouseJSONFile) Load() (w map[int]models.WareHouse, err error) {

	file, err := os.Open(s.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var wareHouseJSON []dto.WareHouseDoc
	err = json.NewDecoder(file).Decode(&wareHouseJSON)
	if err != nil {
		return
	}

	w = make(map[int]models.WareHouse)
	for _, wh := range wareHouseJSON {

		id, err := domain.NewId(wh.Id)
		if err != nil {
			return nil, err
		}

		whCode, err := domain.NewWareHouseCode(wh.WareHouseCode)
		if err != nil {
			return nil, err
		}

		address, err := domain.NewAddress(wh.Address)
		if err != nil {
			return nil, err
		}

		telephone, err := domain.NewTelephone(wh.Telephone)
		if err != nil {
			return nil, err
		}

		minCapacity, err := domain.NewMinimunCapacity(wh.MinimunCapacity)
		if err != nil {
			return nil, err
		}

		minTemperature, err := domain.NewMinimunTemperature(wh.MinimunTemperature)
		if err != nil {
			return nil, err
		}

		w[wh.Id] = models.WareHouse{
			Id:                 id,
			WareHouseCode:      whCode,
			Address:            address,
			Telephone:          telephone,
			MinimunCapacity:    minCapacity,
			MinimunTemperature: minTemperature,
		}
	}

	return
}

func (s *WareHouseJSONFile) Save(wh map[int]models.WareHouse) error {

	// Convert map to list
	wareHouselist := make([]dto.WareHouseDoc, 0, len(wh))
	for _, wareHouse := range wh {
		wareHouselist = append(wareHouselist, dto.WareHouseDoc{
			Id:                 wareHouse.Id.GetId(),
			WareHouseCode:      wareHouse.WareHouseCode.GetWareHouseCode(),
			Address:            wareHouse.Address.GetAddress(),
			Telephone:          wareHouse.Telephone.GetTelephone(),
			MinimunCapacity:    wareHouse.MinimunCapacity.GetMinimunCapacity(),
			MinimunTemperature: wareHouse.MinimunTemperature.GetMinimunTemperature(),
		})
	}

	// Convert list to JSON
	data, err := json.MarshalIndent(wareHouselist, "", " ")
	if err != nil {
		return err
	}

	// Write JSON on file
	err = os.WriteFile(s.path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

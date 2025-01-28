package product

import (
	"encoding/json"
	"errors"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"os"
)

type IProductLoader interface {
	Load() (v map[int]model.Product, err error)
}

type ProductJSONFile struct {
	path string
}

func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{path: path}
}

func (l *ProductJSONFile) Load() (map[int]model.Product, error) {
	file, err := os.Open(l.path)
	if err != nil {
		return nil, errors.New("")
	}

	defer file.Close()

	var ProductJson []model.Product
	err = json.NewDecoder(file).Decode(&ProductJson)
	if err != nil {
		return nil, errors.New("")
	}

	list := make(map[int]model.Product)
	for _, v := range ProductJson {
		list[v.ID] = model.Product{
			ID:            v.ID,
			ProductCode:   v.ProductCode,
			Description:   v.Description,
			ProductTypeID: v.ProductTypeID,
			SellerID:      v.SellerID,
			Dimensions: model.Dimensions{
				Height:    v.Height,
				Length:    v.Length,
				NetWeight: v.NetWeight,
				Width:     v.Width,
			},
			FreezingInfo: model.FreezingInfo{
				ExpirationRate:                 v.ExpirationRate,
				FreezingRate:                   v.FreezingRate,
				RecommendedFreezingTemperature: v.RecommendedFreezingTemperature,
			},
		}
	}

	return list, nil
}

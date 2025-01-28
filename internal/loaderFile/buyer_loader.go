package loaderfile

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/model"
)

type IBuyerLoader interface {
	Load() (v map[int]model.Buyer, err error)
}

type BuyerJSONFile struct {
	path string
}

func NewBuyerJSONFile(path string) *BuyerJSONFile {
	return &BuyerJSONFile{path: path}
}

func (l *BuyerJSONFile) Load() (map[int]model.Buyer, error) {
	file, err := os.Open(l.path)
	if err != nil {
		return nil, errors.New("")
	}

	defer file.Close()

	var buyerJson []model.Buyer
	err = json.NewDecoder(file).Decode(&buyerJson)
	if err != nil {
		return nil, errors.New("")
	}

	list := make(map[int]model.Buyer)
	for _, value := range buyerJson {
		list[value.Id] = model.Buyer{
			Id: value.Id,
			PersonAtributes: model.PersonAtributes{
				Card_Number_Id: value.Card_Number_Id,
				First_Name:     value.First_Name,
				Last_Name:      value.Last_Name,
			},
		}
	}

	return list, nil
}

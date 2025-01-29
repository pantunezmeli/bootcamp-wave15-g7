package loaderfile

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
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

func (l *BuyerJSONFile) Save(buyer model.Buyer) error {
	// Leer el archivo JSON existente
	var buyers []model.Buyer

	file, err := os.Open(l.path)
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		_ = decoder.Decode(&buyers)
	} else if !os.IsNotExist(err) {
		return errors.New("error reading")
	}

	// Agregar el nuevo buyer a la lista
	buyers = append(buyers, buyer)

	file, err = os.Create(l.path)
	if err != nil {
		return errors.New("error writing")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(buyers); err != nil {
		return errors.New("error encoding")
	}

	return nil
}

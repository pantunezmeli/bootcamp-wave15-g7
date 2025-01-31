package buyerstorage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type BuyerJSONFile struct {
	path string
}

func NewBuyerJSONFile(path string) *BuyerJSONFile {
	return &BuyerJSONFile{path: path}
}

func (l *BuyerJSONFile) Load() (map[int]models.Buyer, error) {
	fmt.Println(l.path)
	file, err := os.Open(l.path)
	if err != nil {
		return nil, errors.New("")
	}

	defer file.Close()

	var buyerJson []models.Buyer
	err = json.NewDecoder(file).Decode(&buyerJson)
	if err != nil {
		return nil, errors.New("")
	}

	list := make(map[int]models.Buyer)
	for _, value := range buyerJson {
		list[value.Id] = models.Buyer{
			Id: value.Id,
			PersonAtributes: models.PersonAtributes{
				Card_Number_Id: value.Card_Number_Id,
				First_Name:     value.First_Name,
				Last_Name:      value.Last_Name,
			},
		}
	}

	return list, nil
}

func (l *BuyerJSONFile) Save(buyer models.Buyer) error {
	var buyers []models.Buyer

	file, err := os.Open(l.path)
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		_ = decoder.Decode(&buyers)
	} else if !os.IsNotExist(err) {
		return errors.New("error reading")
	}

	updated := false
	for i := range buyers {
		if buyers[i].Id == buyer.Id {
			buyers[i] = buyer
			updated = true
			break
		}
	}

	if !updated {
		buyers = append(buyers, buyer)
	}

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

func (l *BuyerJSONFile) Delete(buyerID int) error {
	var buyers []models.Buyer

	file, err := os.Open(l.path)
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		_ = decoder.Decode(&buyers)
	} else if !os.IsNotExist(err) {
		return errors.New("error reading")
	}

	var updatedBuyers []models.Buyer
	found := false

	for _, buyer := range buyers {
		if buyer.Id != buyerID {
			updatedBuyers = append(updatedBuyers, buyer)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("buyer not found")
	}

	file, err = os.Create(l.path)
	if err != nil {
		return errors.New("error writing")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(updatedBuyers); err != nil {
		return errors.New("error encoding")
	}

	return nil
}

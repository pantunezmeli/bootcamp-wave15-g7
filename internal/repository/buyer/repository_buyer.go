package buyer

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	loaderfile "github.com/pantunezmeli/bootcamp-wave15-g7/internal/loaderFile"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type BuyerRepository struct {
	buyers map[int]model.Buyer
}

func NewBuyerRepository(buyers map[int]model.Buyer) *BuyerRepository {
	list := make(map[int]model.Buyer)
	if buyers != nil {
		list = buyers
	}
	return &BuyerRepository{buyers: list}
}

func (buyer *BuyerRepository) GetAll() (map[int]model.Buyer, error) {

	list := make(map[int]model.Buyer)
	for key, value := range buyer.buyers {
		list[key] = value
	}

	return list, nil
}

func (buyer *BuyerRepository) GetById(id int) (model.Buyer, error) {

	byer_founded, ok := buyer.buyers[id]
	if !ok {
		return model.Buyer{}, errorbase.ErrNotFound
	}
	return byer_founded, nil
}

func (buyer *BuyerRepository) Create(entity model.Buyer) (model.Buyer, error) {

	_, ok := buyer.buyers[entity.Id]
	if ok {
		return model.Buyer{}, errorbase.ErrConflict
	}

	attributes, err := buyer.ValidateModel(entity)

	if err != nil {
		return model.Buyer{}, err
	}

	entity.Card_Number_Id = attributes["CardNumber"].(domain.CardNumberId).GetCardNumberId()
	entity.First_Name = attributes["FirstName"].(domain.FirstName).GetFirstName()
	entity.Last_Name = attributes["LastName"].(domain.LastName).GetLastName()

	entity.Id = getLastId(buyer.buyers)
	buyer.buyers[entity.Id] = entity

	loader := loaderfile.NewBuyerJSONFile("../docs/db/buyer_data.json")
	loader.Save(entity)

	return entity, nil
}

func getLastId(buyer map[int]model.Buyer) int {
	maxId := 0
	for id := range buyer {
		if id > maxId {
			maxId = id
		}
	}
	return maxId + 1
}

func (*BuyerRepository) ValidateModel(entity model.Buyer) (map[string]any, error) {
	cardNumber, err := domain.NewCardNumberId(entity.Card_Number_Id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	firstName, err := domain.NewFirstName(entity.First_Name)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	lastName, err := domain.NewLastName(entity.Last_Name)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return map[string]any{
		"CardNumber": cardNumber,
		"FirstName":  firstName,
		"LastName":   lastName,
	}, nil
}

func (buyer *BuyerRepository) Update(id int, entity model.Buyer) (model.Buyer, error) {

	element, ok := buyer.buyers[id]
	if !ok {
		return model.Buyer{}, errorbase.ErrConflict
	}
	if cardNumberId, err := domain.NewCardNumberId(id); err == nil {
		element.Card_Number_Id = cardNumberId.GetCardNumberId()
	}

	if firstName, err := domain.NewFirstName(entity.First_Name); err == nil {
		element.First_Name = firstName.GetFirstName()
	}

	if lastName, err := domain.NewLastName(entity.Last_Name); err == nil {
		element.Last_Name = lastName.GetLastName()
	}
	// if err := mergo.Merge(&element, entity, mergo.); err != nil {
	// 	return model.Buyer{}, err
	// }
	loader := loaderfile.NewBuyerJSONFile("../docs/db/buyer_data.json")
	loader.Save(element)
	buyer.buyers[id] = element
	return element, nil
}

//func (buyer *BuyerRepository) Delete(id int) error {

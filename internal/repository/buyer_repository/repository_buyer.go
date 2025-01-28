package buyerrepository

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/model"
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
		return model.Buyer{}, errors.New("not found")
	}
	return byer_founded, nil
}

func (buyer *BuyerRepository) Create(entity model.Buyer) error {

	_, ok := buyer.buyers[entity.Id]
	if !ok {
		return errors.New("element exist")
	}

	entity.Id = getLastId(buyer.buyers)
	buyer.buyers[entity.Id] = entity
	return nil
}

func getLastId(buyer map[int]model.Buyer) int {
	return buyer[len(buyer)-1].Id + 1
}

//func (buyer *BuyerRepository) Update(id int, entity model.Buyer) error {

//func (buyer *BuyerRepository) Delete(id int) error {

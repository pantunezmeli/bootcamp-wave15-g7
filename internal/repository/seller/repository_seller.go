package seller

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	seller_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/seller"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/seller_storage"
)

type SellerStorage struct {
	storage seller_storage.SellerJSONFile 
}

func NewSellerStorage(storage seller_storage.SellerJSONFile) *SellerStorage {
	return &SellerStorage{storage}
}

func (s *SellerStorage) GetAll() (sellers []models.Seller, err error) {
	sellersMap, err := s.storage.Load()
	if err != nil{
		return
	}
	for _, seller := range sellersMap{
		sellers = append(sellers, seller)
	}
	return
}

func  (s *SellerStorage) GetById(id int) (seller models.Seller, err error) {
	sellersMap, err := s.storage.Load()
	if err != nil{
		return
	}
	seller, ok := sellersMap[id]
	if !ok {
		err = ErrSellerNotFound
	}
	return
}

func (s *SellerStorage) Save(modelWithoutId models.Seller) (seller models.Seller, err error) {
	sellersMap, err := s.storage.Load()
	if err != nil {
		return
	}
	if err = s.CheckCid(modelWithoutId, sellersMap); err != nil{
		return
	}
	
	newId := nextId(sellersMap)
	
	id, err := seller_vo.NewSellerId(newId)
	if err != nil {
		return
	} 
	modelWithoutId.ID = id
	seller = modelWithoutId
	sellersMap[newId] = seller

	err = s.storage.Save(sellersMap)
	return
}

func (s *SellerStorage) Delete(id int) (err error){
	sellersMap, err := s.storage.Load()
	if err != nil {
		return
	}
	_, ok := sellersMap[id]
	if !ok {
		err = ErrSellerNotFound
		return
	}

	delete(sellersMap, id)

	err = s.storage.Save(sellersMap)
	return



}

func (s *SellerStorage) CheckCid(sellerModel models.Seller, sellersMap map[int]models.Seller) (err error){
	for _, seller := range sellersMap{
		if seller.Cid == sellerModel.Cid && sellerModel.ID != seller.ID {
			err = ErrCidAlreadyExists
			return
		}
	}
	return
}

func(s *SellerStorage) Update(sellerModel models.Seller) (sellerUpdated models.Seller, err error){
	sellersMap, err := s.storage.Load()
	if err != nil {
		return
	}
	if err = s.CheckCid(sellerModel, sellersMap); err != nil{
		return
	}
	
	sellerUpdated = sellerModel
	sellersMap[int(sellerModel.ID)] = sellerUpdated
	err = s.storage.Save(sellersMap)
	return
}


func nextId(sellersMap map[int]models.Seller) int {
	existingIds := make(map[int]bool)
	for id := range sellersMap {
		existingIds[id] = true
	}

	nextId := 1
	for {
		if !existingIds[nextId] {
			return nextId
		}
		nextId++
	}
}

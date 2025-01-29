package seller

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader"
)

type SellerStorage struct {
	loader loader.SellerJSONFile 
}

func NewSellerStorage(loader loader.SellerJSONFile) *SellerStorage {
	return &SellerStorage{loader}
}

func (s *SellerStorage) GetAll() (sellers []models.Seller, err error) {
	sellersMap, err := s.loader.Load()
	if err != nil{
		return
	}
	for _, seller := range sellersMap{
		sellers = append(sellers, seller)
	}
	return
}

func  (s *SellerStorage) GetById(id int) (seller models.Seller, err error) {
	sellersMap, err := s.loader.Load()
	if err != nil{
		return
	}
	seller, ok := sellersMap[id]
	if !ok {
		err = ErrSellerNotFound
	}
	return
}

func (s *SellerStorage) Save(modelToSave models.Seller) (seller models.Seller, err error) {
	sellersMap, err := s.loader.Load()
	if err != nil {
		return
	}

	if err = s.CheckCid(*modelToSave.Cid.Value(), sellersMap); err != nil{
		return
	}

	if *modelToSave.ID.Value() == 0 {
		newId := nextId(sellersMap)
		
		id, err := domain.NewId(newId)
		if err != nil {
			return models.Seller{}, err
		} 
		modelToSave.ID = id
	}

	seller = modelToSave

	sellersMap[*modelToSave.ID.Value()] = seller

	err = s.loader.Save(sellersMap)

	return

}


func (s *SellerStorage) Delete(id int) (err error){
	sellersMap, err := s.loader.Load()
	if err != nil {
		return
	}
	_, ok := sellersMap[id]
	if !ok {
		err = ErrSellerNotFound
		return
	}

	delete(sellersMap, id)

	err = s.loader.Save(sellersMap)
	return



}

func (s *SellerStorage) CheckCid(cid int, sellersMap map[int]models.Seller) (err error){
	for _, seller := range sellersMap{
		if *seller.Cid.Value() == cid {
			err = ErrCidAlreadyExists
			return
		}
	}
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

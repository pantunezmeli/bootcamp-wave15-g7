package seller

import (
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
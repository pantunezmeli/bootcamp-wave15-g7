package buyerservice

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/model"
	buyerrepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer_repository"
)

type BuyerService struct {
	repo buyerrepository.IRepositoryBuyer
}

func NewBuyerService(rp buyerrepository.IRepositoryBuyer) *BuyerService {
	return &BuyerService{repo: rp}
}

func (repo *BuyerService) GetBuyers() (map[int]model.Buyer, error) {

	list, err := repo.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("empty list")
	}

	return list, nil
}

func (repo *BuyerService) GetBuyer(id int) (model.Buyer, error) {

	if id <= 0 {
		return model.Buyer{}, fmt.Errorf("invalid id")
	}
	buyer, err := repo.repo.GetById(id)
	if err != nil {
		return model.Buyer{}, fmt.Errorf("not founded")
	}
	return buyer, nil
}

func (repo *BuyerService) CreateBuyer(entity model.Buyer) error {
	isValid := entity != model.Buyer{}
	if !isValid {
		return fmt.Errorf("model invalid")
	}

	err := repo.CreateBuyer(entity)
	if err != nil {
		return fmt.Errorf("buyer already exist")
	}

	return nil
}

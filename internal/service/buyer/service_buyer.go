package buyer

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"

	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type BuyerService struct {
	repository buyer.IRepositoryBuyer
}

func NewBuyerService(rp buyer.IRepositoryBuyer) *BuyerService {
	return &BuyerService{repository: rp}
}

func (repo *BuyerService) GetBuyers() ([]dto.BuyerResponse, error) {

	list, err := repo.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("empty list")
	}

	buyers := dto.GenerateResponseList(list)

	return buyers, nil
}

func (repo *BuyerService) GetBuyer(id int) (dto.BuyerResponse, error) {

	if id <= 0 {
		return dto.BuyerResponse{}, fmt.Errorf("invalid id")
	}
	buyer, err := repo.repository.GetById(id)
	if err != nil {
		return dto.BuyerResponse{}, fmt.Errorf("not founded")
	}

	buyerResponse := dto.GenerateBuyerResponse(buyer)
	return buyerResponse, nil
}

func (repo *BuyerService) CreateBuyer(entity model.Buyer) error {
	isValid := entity != model.Buyer{}
	if !isValid {
		return fmt.Errorf("model invalid")
	}

	err := repo.repository.Create(entity)
	if err != nil {
		return fmt.Errorf("buyer already exist")
	}

	return nil
}

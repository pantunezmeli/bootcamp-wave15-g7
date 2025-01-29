package buyer

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"

	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
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
		return nil, errorbase.ErrEmptyList
	}

	buyers := dto.GenerateResponseList(list)
	return buyers, nil
}

func (repo *BuyerService) GetBuyer(id int) (dto.BuyerResponse, error) {

	if id <= 0 {
		return dto.BuyerResponse{}, errorbase.ErrInvalidId
	}
	buyer, err := repo.repository.GetById(id)
	if err != nil {
		return dto.BuyerResponse{}, errorbase.ErrNotFound
	}

	buyerResponse := dto.GenerateBuyerResponse(buyer)
	return buyerResponse, nil
}

func (repo *BuyerService) CreateBuyer(entity model.Buyer) (dto.BuyerResponse, error) {
	buyer, err := repo.repository.Create(entity)
	if err != nil {
		return dto.BuyerResponse{}, errors.New(err.Error())
	}

	buyerResponse := dto.GenerateBuyerResponse(buyer)
	return buyerResponse, nil
}

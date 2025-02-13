package buyer

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"

	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/buyer"
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

func (repo *BuyerService) CreateBuyer(entity models.Buyer) (dto.BuyerResponse, error) {
	buyer, err := repo.repository.Create(entity)

	if err != nil {
		return dto.BuyerResponse{}, err
	}

	buyerResponse := dto.GenerateBuyerResponse(buyer)
	return buyerResponse, nil
}

func (repo *BuyerService) UpdateBuyer(id int, entity dto.BuyerUpdate) (dto.BuyerResponse, error) {

	if id <= 0 {
		return dto.BuyerResponse{}, errorbase.ErrInvalidId
	}

	buyerExist, err2 := repo.GetBuyer(id)
	if err2 != nil {
		return dto.BuyerResponse{}, errorbase.ErrNotFound
	}

	if err3 := dto.ValidateBuyerFields(entity); err3 != nil {
		return dto.BuyerResponse{}, errorbase.ErrUnprocessable
	}

	buyerReq := dto.GenerateBuyerRequestUpdate(id, entity, buyerExist)

	buyer, err := repo.repository.Update(id, buyerReq)
	if err != nil {
		return dto.BuyerResponse{}, err
	}

	buyerResponse := dto.GenerateBuyerResponse(buyer)
	return buyerResponse, nil
}

func (repo *BuyerService) DeleteBuyer(id int) error {
	if id <= 0 {
		return errorbase.ErrInvalidId
	}

	err := repo.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

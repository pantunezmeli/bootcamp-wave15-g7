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

func (service *BuyerService) GetBuyers() ([]dto.BuyerResponse, error) {

	list, err := service.repository.GetAll()
	if err != nil {
		return nil, errorbase.ErrEmptyList
	}

	buyers := dto.GenerateResponseList(list)
	return buyers, nil
}

func (service *BuyerService) GetBuyer(id int) (dto.BuyerResponse, error) {

	if id <= 0 {
		return dto.BuyerResponse{}, errorbase.ErrInvalidId
	}
	buyer, err := service.repository.GetById(id)
	if err != nil {
		return dto.BuyerResponse{}, errorbase.ErrNotFound
	}

	buyerResponse := dto.GenerateBuyerResponse(buyer)
	return buyerResponse, nil
}

func (service *BuyerService) CreateBuyer(entity models.Buyer) (dto.BuyerResponse, error) {
	buyer, err := service.repository.Create(entity)

	if err != nil {
		return dto.BuyerResponse{}, err
	}

	buyerResponse := dto.GenerateBuyerResponse(buyer)
	return buyerResponse, nil
}

func (service *BuyerService) UpdateBuyer(id int, entity dto.BuyerUpdate) (dto.BuyerResponse, error) {

	if id <= 0 {
		return dto.BuyerResponse{}, errorbase.ErrInvalidId
	}

	buyerExist, err2 := service.repository.GetById(id)
	if err2 != nil {
		return dto.BuyerResponse{}, errorbase.ErrNotFound
	}

	if err3 := dto.ValidateBuyerFields(entity); err3 != nil {
		return dto.BuyerResponse{}, errorbase.ErrUnprocessable
	}

	buyerReq := dto.GenerateBuyerRequestUpdate(id, entity, buyerExist)

	buyer, err := service.repository.Update(id, buyerReq)
	if err != nil {
		return dto.BuyerResponse{}, err
	}

	buyerResponse := dto.GenerateBuyerResponse(buyer)
	return buyerResponse, nil
}

func (service *BuyerService) DeleteBuyer(id int) error {
	if id <= 0 {
		return errorbase.ErrInvalidId
	}

	err := service.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

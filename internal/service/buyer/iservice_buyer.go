package buyer

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/buyer"
)

type IServiceBuyer interface {
	GetBuyers() ([]dto.BuyerResponse, error)
	GetBuyer(id int) (dto.BuyerResponse, error)
	CreateBuyer(entity models.Buyer) (dto.BuyerResponse, error)
	UpdateBuyer(id int, entity dto.BuyerResponse) (dto.BuyerResponse, error)
	DeleteBuyer(id int) error
}

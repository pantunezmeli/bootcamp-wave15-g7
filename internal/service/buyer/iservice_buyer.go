package buyer

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type IServiceBuyer interface {
	GetBuyers() ([]dto.BuyerResponse, error)
	GetBuyer(id int) (dto.BuyerResponse, error)
	CreateBuyer(entity model.Buyer) (dto.BuyerResponse, error)
	//UpdateBuyer(id int, entity model.Buyer) error
	//DeleteBuyer(id int) error
}

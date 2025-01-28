package buyerservice

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/model"

type IServiceBuyer interface {
	GetBuyers() (map[int]model.Buyer, error)
	GetBuyer(id int) (model.Buyer, error)
	CreateBuyer(entity model.Buyer) error
	//UpdateBuyer(id int, entity model.Buyer) error
	//DeleteBuyer(id int) error
}

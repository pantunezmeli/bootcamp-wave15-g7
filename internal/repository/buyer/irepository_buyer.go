package buyer

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"

type IRepositoryBuyer interface {
	GetAll() (map[int]model.Buyer, error)
	GetById(id int) (model.Buyer, error)
	Create(entity model.Buyer) error
	//Update(id int, entity model.Buyer) error
	//Delete(id int) error
}

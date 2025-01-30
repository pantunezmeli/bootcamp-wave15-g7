package buyer

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IRepositoryBuyer interface {
	GetAll() (map[int]models.Buyer, error)
	GetById(id int) (models.Buyer, error)
	Create(entity models.Buyer) (models.Buyer, error)
	Update(id int, entity models.Buyer) (models.Buyer, error)
	Delete(id int) error
}

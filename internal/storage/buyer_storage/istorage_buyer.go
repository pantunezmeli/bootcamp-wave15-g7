package buyerstorage

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IBuyerLoader interface {
	Load() (v map[int]models.Buyer, err error)
	Save(buyer models.Buyer) error
	Delete(buyerID int) error
}

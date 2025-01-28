package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
)

type IProductRepository interface {
	GetAll() map[int]model.Product
}

package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
)

type ErrProductRepository struct {
	msg string
}

func (e ErrProductRepository) Error() string {
	return e.msg
}

var (
	ErrProductNotFound = ErrProductRepository{msg: "Product not found"}
)

type IProductRepository interface {
	GetAll() map[int]model.Product
	GetByID(id int) (model.Product, error)
}

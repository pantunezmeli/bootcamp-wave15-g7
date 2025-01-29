package product

import (
	v "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
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
	GetAll() (map[int]model.Product, error)
	GetByID(id int) (model.Product, error)
	CreateProduct(product model.Product) error
	DeleteProduct(id int) error
	ProductCodeExist(productCode string) bool
	GetLastID() v.Id
}

package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
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
	GetAll() (map[int]models.Product, error)
	GetByID(id int) (models.Product, error)
	CreateProduct(product models.Product) error
	DeleteProduct(id int) error
	ProductCodeExist(productCode string) bool
	GetLastID() value_objects.Id
	UpdateProduct(product models.Product) error
}

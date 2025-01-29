package product

import (
	v "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
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
	GetLastID() v.Id
	UpdateProduct(product models.Product) error
}

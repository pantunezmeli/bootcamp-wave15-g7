package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	err "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product/errordb"
)

type ErrProductRepository struct {
	msg string
}

func (e ErrProductRepository) Error() string {
	return e.msg
}

var (
	ErrProductNotFound         = err.ErrNotFound{Message: "Product Not Found"}
	ErrProductCodeAlreadyExist = err.ErrConflict{Message: "Product Code Already Exist"}
)

type IProductRepository interface {
	GetAll() (map[int]models.Product, error)
	GetByID(id int) (models.Product, error)
	CreateProduct(product *models.Product) error
	DeleteProduct(id int) error
	ProductCodeExist(productCode string) error
	UpdateProduct(product models.Product) error
}

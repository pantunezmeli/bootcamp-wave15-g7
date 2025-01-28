package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type IProductService interface {
	GetAll() ([]dto.ProductDTO, error)
	GetByID(id int) (dto.ProductDTO, error)
	DeleteProduct(id int) error
}

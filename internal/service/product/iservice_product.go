package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product"
)

type IProductService interface {
	GetAll() ([]product.ProductDTO, error)
	GetByID(id int) (product.ProductDTO, error)
	DeleteProduct(id int) error
	CreateProduct(product product.ProductDTO) (product.ProductDTO, error)
	UpdateProduct(id int, product product.UpdateProductRequest) (product.ProductDTO, error)
}

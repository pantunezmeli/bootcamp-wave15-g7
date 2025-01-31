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

// ErrService is an interface

type ErrProduct struct {
	message string
}

func (e ErrProduct) Error() string {
	return e.message
}

type ErrValidProduct struct {
	message string
}

func (e ErrValidProduct) Error() string {
	return e.message
}

type ErrNotFoundProduct struct {
	message string
}

func (e ErrNotFoundProduct) Error() string {
	return e.message
}

type ErrProductConflict struct {
	message string
}

func (e ErrProductConflict) Error() string {
	return e.message
}


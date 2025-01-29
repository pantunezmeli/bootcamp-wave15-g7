package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type IProductService interface {
	GetAll() ([]dto.ProductDTO, error)
	GetByID(id int) (dto.ProductDTO, error)
	DeleteProduct(id int) error
	CreateProduct(product dto.ProductDTO) (dto.ProductDTO, error)
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

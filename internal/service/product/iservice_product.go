package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type IProductService interface {
	GetAll() []dto.ProductDTO
	GetByID(id int) (dto.ProductDTO, error)
}

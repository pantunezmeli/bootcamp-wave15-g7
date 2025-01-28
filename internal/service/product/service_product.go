package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

func NewProductService(rp product.IProductRepository) *ProductService {
	return &ProductService{rp: rp}
}

type ProductService struct {
	rp product.IProductRepository
}

func (p ProductService) GetAll() []dto.ProductDTO {
	return dto.ParserListProductToDTO(p.rp.GetAll())
}

func (p ProductService) GetByID(id int) (dto.ProductDTO, error) {
	productSearch, errSearch := p.rp.GetByID(id)
	if errSearch != nil {
		return dto.ProductDTO{}, errSearch
	}

	return dto.ParserProductToDTO(productSearch), nil

}

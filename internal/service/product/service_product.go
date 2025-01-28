package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
)

func NewProductService(rp product.IProductRepository) *ProductService {
	return &ProductService{rp: rp}
}

type ProductService struct {
	rp product.IProductRepository
}

func (p ProductService) GetAll() (map[int]model.Product, error) {
	return p.rp.GetAll(), nil
}

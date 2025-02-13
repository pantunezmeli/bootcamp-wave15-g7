package product_records

import (
	rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/productrecords"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/productrecords"
)

type ProductRecordsService struct {
	rp rp.IProductRecordsRepository
}

func (p ProductRecordsService) CreateProductRecord(newProductRecord dto.ProductRecordsDto) (dto.ProductRecordsDto, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRecordsService) GetProductRecord(productID int) (dto.ProductRecordsDto, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRecordsService(rp rp.IProductRecordsRepository) *ProductRecordsService {
	return &ProductRecordsService{rp: rp}
}

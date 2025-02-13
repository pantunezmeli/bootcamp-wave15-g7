package product_records

import dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/productrecords"

type IProductRecordsService interface {
	CreateProductRecord(newProductRecord dto.ProductRecordsDto) (dto.ProductRecordsDto, error)
	GetProductRecord(productID int) (dto.ProductRecordsDto, error)
}

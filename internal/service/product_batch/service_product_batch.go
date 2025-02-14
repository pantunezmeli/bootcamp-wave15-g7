package productbatch

import (
	productbatch "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product_batch"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product_batch"
)

// ProductService is a struct that represents a ProductBatch service
type ProductBatchService struct {
	repository productbatch.IProductBatchRepository
}

// NewProductService is a function that returns a new instance of ProductService
func NewProductBatchService(repo productbatch.IProductBatchRepository) *ProductBatchService {
	return &ProductBatchService{repository: repo}
}

// Create is a method that creates a new ProductBatch
func (s *ProductBatchService) CreateProductBatch(productBatchRequest dto.ProductBatchResponse) (dto.ProductBatchResponse, error) {

	productBatch, err := dto.GenerateProductBatchDtoToModel(productBatchRequest)
	if err != nil {
		return dto.ProductBatchResponse{}, err
	}

	batchStored, err := s.repository.Store(productBatch)
	if err != nil {
		return dto.ProductBatchResponse{}, err

	}

	return dto.GenerateProductBatchModelToDto(batchStored), nil
}

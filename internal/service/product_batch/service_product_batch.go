package productbatch

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
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
func (s *ProductBatchService) CreateProductBatch(productBatch models.ProductBatch) (dto.ProductBatchResponse, error) {
	return s.repository.Store(&productBatch)
}

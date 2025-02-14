package productbatch

import (
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product_batch"
)

// SectionService is an interface that represents a section service
type IProductBatchService interface {
	// CreateSection is a method that creates a new Section
	CreateProductBatch(v dto.ProductBatchResponse) (dto.ProductBatchResponse, error)
}

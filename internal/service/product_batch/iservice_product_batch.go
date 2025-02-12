package productbatch

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product_batch"
)

// SectionService is an interface that represents a section service
type IProductBatchService interface {
	// CreateSection is a method that creates a new Section
	CreateProductBatch(v models.ProductBatch) (dto.ProductBatchResponse, error)
}

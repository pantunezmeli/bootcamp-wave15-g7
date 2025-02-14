package productbatch

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

// IProductBatchRepository is an interface that represents a product batch repository
type IProductBatchRepository interface {
	// Create is a method that creates a new product batch
	Store(v models.ProductBatch) (models.ProductBatch, error)
}

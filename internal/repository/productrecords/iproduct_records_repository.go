package productrecords

import (
	m "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product/errordb"
)

var (
	ErrRecordsNotFound = errordb.ErrNotFound{Message: "Records Not Found"}
)

type IProductRecordsRepository interface {
	CreateProductRecord(newProductRecord *m.ProductRecords) error
	GetRecordsDataOptionalId(id *int) ([]m.RecordsData, error)
}

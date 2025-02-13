package productrecords

import m "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IProductRecordsRepository interface {
	CreateProductRecord(newProductRecord *m.ProductRecords) error
	GetById(id int) (m.ProductRecords, error)
}

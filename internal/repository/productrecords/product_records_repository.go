package productrecords

import (
	"database/sql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type ProductRecordsRepository struct {
	DB *sql.DB
}

func NewProductRecordsRepository(db *sql.DB) *ProductRecordsRepository {
	return &ProductRecordsRepository{DB: db}
}

func (p ProductRecordsRepository) CreateProductRecord(newProductRecord *models.ProductRecords) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRecordsRepository) GetById(id int) (models.ProductRecords, error) {
	//TODO implement me
	panic("implement me")
}

package productrecords

import (
	"database/sql"
	"errors"
	m "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	errdb "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product/errordb"
)

type ProductRecordsRepository struct {
	db *sql.DB
}

func NewProductRecordsRepository(db *sql.DB) *ProductRecordsRepository {
	return &ProductRecordsRepository{db: db}
}

func (r ProductRecordsRepository) CreateProductRecord(newProductRecord *m.ProductRecords) error {
	//TODO implement me
	panic("implement me")
}

func (r ProductRecordsRepository) GetRecordsDataOptionalId(id *int) (records []m.RecordsData, err error) {
	QUERY :=
		`
		 select p.id, p.description, count(pr.id) as records
		 from products p
		 inner join product_records pr on pr.product_id = p.id
		 where (p.id = ? OR ? IS NULL)
		 group by p.id
		 `
	rows, errQuery := r.db.Query(QUERY, id, id)
	if errQuery != nil {
		err = errdb.ErrDB{Message: "Error getAll Records Data"}
		return
	}
	defer rows.Close()

	if errMap := r.getAllEntity(rows, &records); errMap != nil {
		if errors.Is(errMap, ErrRecordsNotFound) {
			err = errMap
			return
		} else {
			err = errdb.ErrDB{Message: "Error mapping Records Data"}
			return
		}
	}

	return
}

func (r ProductRecordsRepository) getAllEntity(rows *sql.Rows, records *[]m.RecordsData) (err error) {
	for rows.Next() {
		var record m.RecordsData
		if errMap := r.getEntity(rows, &record); errMap != nil {
			err = errMap
			return
		}
		*records = append(*records, record)
	}

	if len(*records) == 0 {
		err = ErrRecordsNotFound
		return
	}

	return nil
}

func (r ProductRecordsRepository) getEntity(rows *sql.Rows, record *m.RecordsData) (err error) {
	var idProduct int

	if errScan := rows.Scan(&idProduct, &record.Description, &record.RecordsCount); errScan != nil {
		return errdb.ErrDB{Message: "Error reading Records Data"}
	}

	if errRows := rows.Err(); errRows != nil {
		err = errdb.ErrDB{Message: "Error map Records Data"}
		return
	}

	record.ProductId, err = value_objects.NewId(idProduct)
	return
}

package productrecords

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	m "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	errdb "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product/errordb"
	"strings"
)

type ProductRecordsRepository struct {
	db *sql.DB
}

const FK_PRODYCT_ID = "product_id"

func NewProductRecordsRepository(db *sql.DB) *ProductRecordsRepository {
	return &ProductRecordsRepository{db: db}
}

func (r ProductRecordsRepository) CreateProductRecord(newRecord *m.ProductRecords) (err error) {

	result, errQuery := r.db.Exec(`insert into product_records (last_update_date, purchase_price, sale_price, product_id) values (?, ?, ?, ?)`,
		newRecord.LastUpdateDate, newRecord.PurchasePrice, newRecord.SalePrice, newRecord.ProductId.GetId())
	if errQuery != nil {
		err = errQuery
		r.errorMysql(&err, "Error creating product")
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = errdb.ErrDB{Message: "Error getting last insert id"}
		return
	}
	newRecord.Id, err = value_objects.NewId(int(id))
	if err != nil {
		return
	}

	return

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

func (r ProductRecordsRepository) errorMysql(err *error, message string) {
	var mysqlErr *mysql.MySQLError
	if errors.As(*err, &mysqlErr) {
		switch mysqlErr.Number {

		case 1452:
			r.getFKError(message, mysqlErr, err)
			return
		}
	}
	*err = errdb.ErrDB{Message: message}
}

func (r ProductRecordsRepository) getFKError(message string, mysqlErr *mysql.MySQLError, err *error) {
	switch {

	case strings.Contains(mysqlErr.Message, FK_PRODYCT_ID):
		*err = errdb.ErrViolateFK{Message: "Product not found"}

	default:
		*err = errdb.ErrDB{Message: message}

	}
}

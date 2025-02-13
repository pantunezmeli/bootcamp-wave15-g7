package product

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	m "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	errdb "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product/errordb"
	"strings"
)

const (
	FK_PRODYCT_TYPE = "product_type_id"
	FK_SELLER       = "seller_id"
)

type RepositoryProductMysql struct {
	db *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *RepositoryProductMysql {
	return &RepositoryProductMysql{db: db}
}

func (r RepositoryProductMysql) GetAll() (products map[int]m.Product, err error) {
	QUERY := `select id, description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id from products`

	rows, errQuery := r.db.Query(QUERY)
	if errQuery != nil {
		err = errdb.ErrDB{Message: "Error getAll products"}
		return
	}
	defer rows.Close()

	products = make(map[int]m.Product)
	if errMap := r.getAllEntity(rows, &products); errMap != nil {
		if errors.Is(errMap, ErrProductNotFound) {
			err = errMap
			return
		} else {
			err = errdb.ErrDB{Message: "Error mapping products"}
			return
		}
	}

	return
}

func (r RepositoryProductMysql) GetByID(id int) (product m.Product, err error) {
	QUERY := `select id, description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id from products where id = ?`

	rows, errQuery := r.db.Query(QUERY, id)
	if errQuery != nil {
		err = errdb.ErrDB{Message: "Error querying GetById product"}
		return
	}

	if rows.Next() {
		if errMap := r.getEntity(rows, &product); errMap != nil {
			if errors.Is(errMap, ErrProductNotFound) {
				err = errMap
				return
			} else {
				err = errdb.ErrDB{Message: "Error mapping product"}
				return
			}
		}
	} else {
		err = ErrProductNotFound
		return
	}

	return
}

func (r RepositoryProductMysql) CreateProduct(product *m.Product) (err error) {
	result, errQuery := r.db.Exec(`insert into products (description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		product.Description, product.ExpirationRate, product.FreezingRate, product.Height, product.Length, product.NetWeight, product.ProductCode, product.RecommendedFreezingTemperature, product.Width, product.ProductTypeID, product.SellerID.Value())
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
	product.ID, err = value_objects.NewProductId(int(id))
	if err != nil {
		return
	}

	return
}

func (r RepositoryProductMysql) DeleteProduct(id int) (err error) {
	_, err = r.db.Exec(`DELETE FROM products WHERE id = ?`, id)
	if err != nil {
		r.errorMysql(&err, "Error deleting product")
	}
	return
}

func (r RepositoryProductMysql) ProductCodeExist(productCode string) (err error) {
	QUERY := `SELECT count(1) FROM products WHERE product_code = ?`

	var count int
	err = r.db.QueryRow(QUERY, productCode).Scan(&count)
	if err != nil {
		err = errdb.ErrDB{Message: "Error valid product code"}
		return
	}

	if count > 0 {
		err = ErrProductCodeAlreadyExist
		return
	}

	return
}

func (r RepositoryProductMysql) UpdateProduct(product m.Product) (err error) {
	_, errQuery := r.db.Exec(`update products set  description = ?, expiration_rate = ?, freezing_rate = ?, height = ?, length = ?, net_weight = ?, product_code = ?, recommended_freezing_temperature = ?, width = ?, product_type_id = ?, seller_id = ? where id = ?`,
		product.Description, product.ExpirationRate, product.FreezingRate, product.Height, product.Length, product.NetWeight, product.ProductCode, product.RecommendedFreezingTemperature, product.Width, product.ProductTypeID, product.SellerID.Value(), product.ID)
	if errQuery != nil {
		err = errQuery
		r.errorMysql(&err, "Error update product")
		return
	}
	return
}

func (r RepositoryProductMysql) getAllEntity(rows *sql.Rows, products *map[int]m.Product) (err error) {
	for rows.Next() {
		var product m.Product
		if errMap := r.getEntity(rows, &product); errMap != nil {
			err = errMap
			return
		}
		key := int(product.ID)
		(*products)[key] = product
	}

	if len(*products) == 0 {
		err = ErrProductNotFound
		return
	}

	return nil
}

func (r RepositoryProductMysql) getEntity(rows *sql.Rows, p *m.Product) (err error) {
	var idSeller int

	if errScan := rows.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.NetWeight, &p.ProductCode, &p.RecommendedFreezingTemperature, &p.Width, &p.ProductTypeID, &idSeller); errScan != nil {
		return errdb.ErrDB{Message: "Error reading product"}
	}

	if errRows := rows.Err(); errRows != nil {
		err = errdb.ErrDB{Message: "Error map product"}
		return
	}

	p.SellerID, err = value_objects.NewSellerId(idSeller)
	return
}

func (r RepositoryProductMysql) errorMysql(err *error, message string) {
	var mysqlErr *mysql.MySQLError
	if errors.As(*err, &mysqlErr) {
		switch mysqlErr.Number {

		case 1451:
			*err = errdb.ErrConflict{Message: message}
			return

		case 1452:
			r.getFKError(message, mysqlErr, err)
			return
		}
	}
	*err = errdb.ErrDB{Message: message}
}

func (r RepositoryProductMysql) getFKError(message string, mysqlErr *mysql.MySQLError, err *error) {
	switch {

	case strings.Contains(mysqlErr.Message, FK_PRODYCT_TYPE):
		*err = errdb.ErrViolateFK{Message: "Product Type not found"}

	case strings.Contains(mysqlErr.Message, FK_SELLER):
		*err = errdb.ErrViolateFK{Message: "Seller not found"}

	default:
		*err = errdb.ErrDB{Message: message}

	}
}

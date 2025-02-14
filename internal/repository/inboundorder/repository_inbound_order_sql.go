package inboundorder

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

func NewInboundOrderSQL(database *sql.DB) *InboundOrderSQL {
	return &InboundOrderSQL{db: database}
}

type InboundOrderSQL struct {
	db *sql.DB
}

func (r *InboundOrderSQL) New(inboundOrder models.InboundOrder) (newInboundOrder models.InboundOrder, err error) {
	query := "INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, inboundOrder.OrderDate.GetDate(), inboundOrder.OrderNumber.GetOrderNumber(), inboundOrder.EmployeeId.GetId(), inboundOrder.ProductBatchId.GetId(), inboundOrder.WareHouseId.GetId())
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok {
			if mysqlErr.Number == 1062 { // Duplicate entry (order_number is the only field that doesn't accept duplicates)
				err = ErrOrderNumberNotUnique
				return
			}
			if mysqlErr.Number == 1452 { // A foregin key constraint fails
				if strings.Contains(mysqlErr.Message, "employees") {
					err = ErrEmployeeIdNotFound
				}
				if strings.Contains(mysqlErr.Message, "product_batches") {
					err = ErrProductBatchIdNotFound
				}
				if strings.Contains(mysqlErr.Message, "warehouses") {
					err = ErrWarehouseIdNotFound
				}
				return
			}
		}
		err = ErrDatabase
		return
	}

	newInboundOrder = inboundOrder
	id64, err := result.LastInsertId()
	if err != nil {
		err = ErrDatabase
		return
	}
	newInboundOrder.Id = value_objects.NewOptionalId(int(id64))

	return
}

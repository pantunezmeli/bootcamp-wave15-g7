package purchase

import (
	"database/sql"

	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/purchase"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/querys"
)

type PurchaseRepository struct {
	db *sql.DB
}

func NewBuyerRepository(db *sql.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (repo *PurchaseRepository) GetReportPurchaseOrders() ([]dto.ReportPurchaseOrders, error) {
	rows, err := repo.db.Query(querys.GetReportPurchaseOrders)
	if err != nil {
		return nil, errorbase.ErrDatabaseOperationFailed
	}
	defer rows.Close()

	var purchases []dto.ReportPurchaseOrders
	for rows.Next() {
		var purchase dto.ReportPurchaseOrders
		err := rows.Scan(&purchase.ID, &purchase.CardNumberID, &purchase.FirstName, &purchase.LastName, &purchase.PurchaseOrdersCount)
		if err != nil {
			return nil, errorbase.ErrDatabaseOperationFailed
		}
		purchases = append(purchases, purchase)
	}

	if err = rows.Err(); err != nil {
		return nil, errorbase.ErrDatabaseOperationFailed
	}
	return purchases, nil
}

func (repo *PurchaseRepository) GetReportPurchaseOrdersById(id int) (dto.ReportPurchaseOrders, error) {
	if exist := repo.buyerExist(id); !exist {
		return dto.ReportPurchaseOrders{}, errorbase.ErrNotFound
	}
	row := repo.db.QueryRow(querys.GetReportPurchaseOrdersById, id)
	var purchase dto.ReportPurchaseOrders

	err := row.Scan(&purchase.ID, &purchase.CardNumberID, &purchase.FirstName, &purchase.LastName, &purchase.PurchaseOrdersCount)

	if err != nil {
		return dto.ReportPurchaseOrders{}, errorbase.ErrDatabaseOperationFailed
	}

	return purchase, nil
}

func (repo *PurchaseRepository) buyerExist(id int) bool {
	var exist bool
	err := repo.db.QueryRow(querys.ExistsBuyer, id).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

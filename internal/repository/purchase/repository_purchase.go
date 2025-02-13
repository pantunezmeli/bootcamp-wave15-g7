package purchase

import (
	"database/sql"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/querys"
)

type PurchaseRepository struct {
	db *sql.DB
}

func NewBuyerRepository(db *sql.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (repo *PurchaseRepository) GetReportPurchaseOrders() ([]models.ReportPurchaseOrders, error) {
	rows, err := repo.db.Query(querys.GetReportPurchaseOrders)
	if err != nil {
		return nil, errorbase.ErrDatabaseOperationFailed
	}
	defer rows.Close()

	var purchases []models.ReportPurchaseOrders
	for rows.Next() {
		var purchase models.ReportPurchaseOrders
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

func (repo *PurchaseRepository) GetReportPurchaseOrdersById(id int) (models.ReportPurchaseOrders, error) {
	if exist := repo.buyerExist(id); !exist {
		return models.ReportPurchaseOrders{}, errorbase.ErrNotFound
	}
	row := repo.db.QueryRow(querys.GetReportPurchaseOrdersById, id)
	var purchase models.ReportPurchaseOrders

	err := row.Scan(&purchase.ID, &purchase.CardNumberID, &purchase.FirstName, &purchase.LastName, &purchase.PurchaseOrdersCount)

	if err != nil {
		return models.ReportPurchaseOrders{}, errorbase.ErrDatabaseOperationFailed
	}

	return purchase, nil
}

func (repo *PurchaseRepository) CreatePurchaseOrder(entity models.PurchaseOrder) (models.PurchaseOrder, error) {

	attributes, err := repo.validateModel(entity)
	if err != nil {
		return models.PurchaseOrder{}, err
	}

	order_number := attributes["Order_number"].(value_objects.OrderNumber).GetOrderNumber()
	order_date := attributes["Order_date"].(value_objects.OrderDate).GetOrderDate()
	tracking_code := attributes["Tracking_code"].(value_objects.TrackingCode).GetTrackingCode()
	buyerId := attributes["Buyer_ID"].(value_objects.BuyerID).GetBuyerID()
	carrierId := attributes["Carrier_ID"].(value_objects.CarrierID).GetCarrierID()
	order_status_id := attributes["Order_Status_ID"].(value_objects.OrderStatusID).GetOrderStatusID()
	warehouseId := attributes["Warehouse_ID"].(value_objects.WarehouseID).GetWarehouseID()

	validationChecks := map[bool]error{
		!repo.buyerExist(entity.Buyer_ID):              errorbase.ErrBuyerFKNotExist,
		!repo.carrierExist(entity.Carrier_ID):          errorbase.ErrCarrierFKNotExist,
		!repo.orderStatusExist(entity.Order_Status_ID): errorbase.ErrOrderStatusFKNotExist,
		!repo.warehouseExist(entity.Warehouse_ID):      errorbase.ErrWareHouseFKNotExist,
	}

	for condition, err := range validationChecks {
		if condition {
			return models.PurchaseOrder{}, err
		}
	}

	result, err := repo.db.Exec(querys.CreatePurchaseOrder, order_number, order_date, tracking_code, buyerId, carrierId, order_status_id, warehouseId)
	if err != nil {
		return models.PurchaseOrder{}, errorbase.ErrDatabaseOperationFailed
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.PurchaseOrder{}, errorbase.ErrDatabaseOperationFailed
	}

	entity.ID = int(id)

	return entity, nil

}

func (*PurchaseRepository) validateModel(entity models.PurchaseOrder) (map[string]any, error) {

	order_number, err := value_objects.NewOrderNumber(entity.Order_number)
	if err != nil {
		return nil, err
	}

	order_date, err := value_objects.NewOrderDate(entity.Order_date)
	if err != nil {
		return nil, err
	}

	tracking_code, err := value_objects.NewTrackingCode(entity.Tracking_code)
	if err != nil {
		return nil, err
	}

	buyerId, err := value_objects.NewBuyerID(entity.Buyer_ID)
	if err != nil {
		return nil, err
	}

	carrierId, err := value_objects.NewCarrierID(entity.Carrier_ID)
	if err != nil {
		return nil, err
	}

	order_status_id, err := value_objects.NewOrderStatusID(entity.Order_Status_ID)
	if err != nil {
		return nil, err
	}

	warehouseId, err := value_objects.NewWarehouseID(entity.Warehouse_ID)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"Order_number":    order_number,
		"Order_date":      order_date,
		"Tracking_code":   tracking_code,
		"Buyer_ID":        buyerId,
		"Carrier_ID":      carrierId,
		"Order_Status_ID": order_status_id,
		"Warehouse_ID":    warehouseId,
	}, nil
}

func (repo *PurchaseRepository) buyerExist(id int) bool {
	var exist bool
	err := repo.db.QueryRow(querys.ExistsBuyer, id).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

func (repo *PurchaseRepository) carrierExist(id int) bool {
	var exist bool
	err := repo.db.QueryRow(querys.CarrierExist, id).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

func (repo *PurchaseRepository) orderStatusExist(id int) bool {
	var exist bool
	err := repo.db.QueryRow(querys.OrderStatusExist, id).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}
func (repo *PurchaseRepository) warehouseExist(id int) bool {
	var exist bool
	err := repo.db.QueryRow(querys.Warehouse, id).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

// BL validation
func (repo *PurchaseRepository) OrderNumberExist(order string) bool {
	var exist bool
	err := repo.db.QueryRow(querys.OrderNumberExist, order).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}
func (repo *PurchaseRepository) TrackingCodeExist(code string) bool {
	var exist bool
	err := repo.db.QueryRow(querys.TrackingCodeExist, code).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

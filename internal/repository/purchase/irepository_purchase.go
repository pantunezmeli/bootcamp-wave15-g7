package purchase

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type IRepositoryPurchase interface {
	GetReportPurchaseOrders() ([]models.ReportPurchaseOrders, error)
	GetReportPurchaseOrdersById(id int) (models.ReportPurchaseOrders, error)
	CreatePurchaseOrder(entity models.PurchaseOrder) (models.PurchaseOrder, error)
	ElementExist(number string, query string) bool
}

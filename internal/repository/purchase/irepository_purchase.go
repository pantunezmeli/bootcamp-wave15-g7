package purchase

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/purchase"
)

type IRepositoryPurchase interface {
	GetReportPurchaseOrders() ([]dto.ReportPurchaseOrders, error)
	GetReportPurchaseOrdersById(id int) (dto.ReportPurchaseOrders, error)
	CreatePurchaseOrder(entity models.PurchaseOrder) (models.PurchaseOrder, error)
	OrderNumberExist(order string) bool
	TrackingCodeExist(code string) bool
}

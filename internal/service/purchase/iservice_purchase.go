package purchase

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/purchase"
)

type IServicePurchase interface {
	GetReport() ([]models.ReportPurchaseOrders, error)
	GetReportById(id int) (models.ReportPurchaseOrders, error)
	CreatePurchase(entity dto.PurchaseOrderResponse) (dto.PurchaseOrderResponse, error)
}

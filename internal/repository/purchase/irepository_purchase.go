package purchase

import (
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/purchase"
)

type IRepositoryPurchase interface {
	GetReportPurchaseOrders() ([]dto.ReportPurchaseOrders, error)
	GetReportPurchaseOrdersById(id int) (dto.ReportPurchaseOrders, error)
}

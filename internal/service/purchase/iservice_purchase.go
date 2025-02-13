package purchase

import (
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/purchase"
)

type IServicePurchase interface {
	GetReport() ([]dto.ReportPurchaseOrders, error)
	GetReportById(id int) (dto.ReportPurchaseOrders, error)
}

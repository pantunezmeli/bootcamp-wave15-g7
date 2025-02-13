package purchase

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/purchase"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/purchase"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type PurchaseService struct {
	repository purchase.IRepositoryPurchase
}

func NewBuyerService(rp purchase.IRepositoryPurchase) *PurchaseService {
	return &PurchaseService{repository: rp}
}

func (service *PurchaseService) GetReport() ([]dto.ReportPurchaseOrders, error) {

	list, err := service.repository.GetReportPurchaseOrders()
	if errors.Is(err, errorbase.ErrDatabaseOperationFailed) {
		return nil, errorbase.ErrDatabaseOperationFailed
	} else if err != nil {
		return nil, errorbase.ErrEmptyList
	}

	return list, nil
}

func (service *PurchaseService) GetReportById(id int) (dto.ReportPurchaseOrders, error) {
	if id <= 0 {
		return dto.ReportPurchaseOrders{}, errorbase.ErrInvalidId
	}
	reportPurchase, err := service.repository.GetReportPurchaseOrdersById(id)
	if errors.Is(err, errorbase.ErrNotFound) {
		return dto.ReportPurchaseOrders{}, errorbase.ErrNotFound
	} else if err != nil {
		return dto.ReportPurchaseOrders{}, errorbase.ErrDatabaseOperationFailed
	}
	return reportPurchase, nil
}

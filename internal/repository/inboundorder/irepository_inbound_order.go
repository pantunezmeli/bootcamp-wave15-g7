package inboundorder

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

var ErrDatabase = errors.New("database error")
var ErrEmployeeIdNotFound = errors.New("employee not found")
var ErrProductBatchIdNotFound = errors.New("product batch not found")
var ErrWarehouseIdNotFound = errors.New("warehouse not found")
var ErrOrderNumberNotUnique = errors.New("order number must be unique")

type InboundOrderRepository interface {
	New(inboundOrder models.InboundOrder) (newInboundOrder models.InboundOrder, err error)
}

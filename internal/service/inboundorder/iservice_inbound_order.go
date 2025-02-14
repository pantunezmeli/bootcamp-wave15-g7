package inboundorder

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/inboundorder"
)

var (
	ErrEmployeeNotFound         = errors.New("employee not found")
	ErrProductBatchNotFound     = errors.New("product batch not found")
	ErrWarehouseNotFound        = errors.New("warehouse not found")
	ErrEmptyField               = errors.New("employee data lacks a required field")
	ErrOrderNumberAlreadyExists = errors.New("employee order number already exists")
)

type InboundOrderService interface {
	New(inboundOrderData inboundorder.InboundOrderDoc) (newInboundOrderData inboundorder.InboundOrderDoc, err error)
}

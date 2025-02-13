package models

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"

type InboundOrder struct {
	Id             value_objects.Id
	OrderDate      value_objects.ValidatedDate
	OrderNumber    value_objects.OrderNumber
	EmployeeId     value_objects.Id
	ProductBatchId value_objects.Id
	WareHouseId    value_objects.Id
}

package inboundorder

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type InboundOrderDoc struct {
	Id             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId int    `json:"product_batch_id"`
	WareHouseId    int    `json:"warehouse_id"`
}

func InboundOrderDtoToModel(dto InboundOrderDoc) (inboundOrder models.InboundOrder, err error) {
	newId, errValidation := value_objects.NewId(dto.Id)
	if errValidation != nil {
		return
	}
	newOrderDate, errValidation := value_objects.NewDate(dto.OrderDate)
	if errValidation != nil {
		return
	}
	newOrderNumber, errValidation := value_objects.NewOrderNumber(dto.OrderNumber)
	if errValidation != nil {
		return
	}
	newEmployeeId, errValidation := value_objects.NewId(dto.EmployeeId)
	if errValidation != nil {
		return
	}
	newProductBatchId, errValidation := value_objects.NewId(dto.ProductBatchId)
	if errValidation != nil {
		return
	}
	newWareHouseId, errValidation := value_objects.NewId(dto.WareHouseId)
	if errValidation != nil {
		return
	}
	inboundOrder = models.InboundOrder{
		Id:             newId,
		OrderDate:      newOrderDate,
		OrderNumber:    newOrderNumber,
		EmployeeId:     newEmployeeId,
		ProductBatchId: newProductBatchId,
		WareHouseId:    newWareHouseId,
	}
	return
}

func InboundOrderModelToDto(inboundOrder models.InboundOrder) (dto InboundOrderDoc) {
	dto = InboundOrderDoc{
		Id:             inboundOrder.Id.GetId(),
		OrderDate:      inboundOrder.OrderDate.GetDate(),
		OrderNumber:    inboundOrder.OrderNumber.GetOrderNumber(),
		EmployeeId:     inboundOrder.EmployeeId.GetId(),
		ProductBatchId: inboundOrder.ProductBatchId.GetId(),
		WareHouseId:    inboundOrder.WareHouseId.GetId(),
	}
	return
}

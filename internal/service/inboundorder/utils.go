package inboundorder

import dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/inboundorder"

func validateFields(inboundOrderData dto.InboundOrderDoc) (err error) {
	if inboundOrderData.OrderNumber == "" || inboundOrderData.OrderDate == "" || inboundOrderData.EmployeeId == 0 || inboundOrderData.ProductBatchId == 0 || inboundOrderData.WareHouseId == 0 {
		err = ErrEmptyField
	}
	return
}

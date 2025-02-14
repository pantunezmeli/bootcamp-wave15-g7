package dto

import (
	"sort"
	"time"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type PurchaseOrderResponse struct {
	ID            int    `json:"id"`
	OrderNumber   string `json:"order_number"`
	OrderDate     string `json:"order_date"`
	TrackingCode  string `json:"tracking_code"`
	BuyerID       int    `json:"buyer_id" validate:"required"`
	CarrierID     int    `json:"carrier_id" validate:"required"`
	OrderStatusID int    `json:"order_status_id" validate:"required"`
	WarehouseID   int    `json:"warehouse_id" validate:"required"`
}

func GeneratePurchaseResponseList(purchases []models.PurchaseOrder) []PurchaseOrderResponse {
	var list []PurchaseOrderResponse
	for _, value := range purchases {
		list = append(list, PurchaseOrderResponse{
			ID:            value.ID,
			OrderNumber:   value.Order_number,
			OrderDate:     value.Order_date.Format("2006-01-02"),
			TrackingCode:  value.Tracking_code,
			BuyerID:       value.Buyer_ID,
			CarrierID:     value.Carrier_ID,
			OrderStatusID: value.Order_Status_ID,
			WarehouseID:   value.Warehouse_ID,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}

func GeneratePurchaseResponse(new models.PurchaseOrder) PurchaseOrderResponse {

	purchase := PurchaseOrderResponse{
		ID:            new.ID,
		OrderNumber:   new.Order_number,
		OrderDate:     new.Order_date.Format("2006-01-02"),
		TrackingCode:  new.Tracking_code,
		BuyerID:       new.Buyer_ID,
		CarrierID:     new.Carrier_ID,
		OrderStatusID: new.Order_Status_ID,
		WarehouseID:   new.Warehouse_ID,
	}
	return purchase
}

func GeneratePurchaseResponseToEntity(new PurchaseOrderResponse) (models.PurchaseOrder, error) {

	parsedTime, err := time.Parse("2006-01-02", new.OrderDate)
	if err != nil {
		return models.PurchaseOrder{}, errorbase.ErrInvalidRequest
	}

	purchase := models.PurchaseOrder{
		ID:              new.ID,
		Order_number:    new.OrderNumber,
		Order_date:      parsedTime,
		Tracking_code:   new.TrackingCode,
		Buyer_ID:        new.BuyerID,
		Carrier_ID:      new.CarrierID,
		Order_Status_ID: new.OrderStatusID,
		Warehouse_ID:    new.WarehouseID,
	}
	return purchase, nil
}

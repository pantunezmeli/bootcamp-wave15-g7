package dto

import (
	"sort"
	"time"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type PurchaseOrderResponse struct {
	ID            int       `json:"id"`
	OrderNumber   string    `json:"order_number"`
	OrderDate     time.Time `json:"order_date"`
	TrackingCode  string    `json:"tracking_code"`
	BuyerID       int       `json:"buyer_id" validate:"required"`
	CarrierID     int       `json:"carrier_id" validate:"required"`
	OrderStatusID int       `json:"order_status_id" validate:"required"`
	WarehouseID   int       `json:"warehouse_id" validate:"required"`
}

func GeneratePurchaseResponseList(buyers []models.PurchaseOrder) []PurchaseOrderResponse {
	var list []PurchaseOrderResponse
	for _, value := range buyers {
		list = append(list, PurchaseOrderResponse{
			ID:            value.ID,
			OrderNumber:   value.Order_number,
			OrderDate:     value.Order_date,
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

type ReportPurchaseOrders struct {
	ID                  int    `json:"id"`
	CardNumberID        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

package value_objects

import (
	"strings"
	"time"

	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type OrderNumber struct {
	value string
}

func NewOrderNumber(orderNumber string) (OrderNumber, error) {
	orderNumber = strings.TrimSpace(orderNumber)
	if orderNumber == "" {
		return OrderNumber{}, errorbase.ErrEmptyParameters
	}
	return OrderNumber{value: orderNumber}, nil
}

func (o OrderNumber) GetOrderNumber() string {
	return o.value
}

type OrderDate struct {
	value time.Time
}

func NewOrderDate(orderDate time.Time) (OrderDate, error) {
	formattedDate := orderDate.Format("2006-01-02")

	parsedDate, err := time.Parse("2006-01-02", formattedDate)
	if err != nil {
		return OrderDate{}, errorbase.ErrInvalidRequest
	}

	return OrderDate{value: parsedDate}, nil
}

func (o OrderDate) GetOrderDate() time.Time {
	return o.value
}

type TrackingCode struct {
	value string
}

func NewTrackingCode(trackingCode string) (TrackingCode, error) {
	trackingCode = strings.TrimSpace(trackingCode)
	if trackingCode == "" {
		return TrackingCode{}, errorbase.ErrEmptyParameters
	}
	return TrackingCode{value: trackingCode}, nil
}

func (t TrackingCode) GetTrackingCode() string {
	return t.value
}

type BuyerID struct {
	value int
}

func NewBuyerID(buyerID int) (BuyerID, error) {

	if buyerID <= 0 {
		return BuyerID{}, errorbase.ErrInvalidIdField
	}
	return BuyerID{value: buyerID}, nil
}

func (b BuyerID) GetBuyerID() int {
	return b.value
}

type CarrierID struct {
	value int
}

func NewCarrierID(carrierID int) (CarrierID, error) {
	if carrierID <= 0 {
		return CarrierID{}, errorbase.ErrInvalidIdField
	}
	return CarrierID{value: carrierID}, nil
}

func (c CarrierID) GetCarrierID() int {
	return c.value
}

type OrderStatusID struct {
	value int
}

func NewOrderStatusID(orderStatusID int) (OrderStatusID, error) {
	if orderStatusID <= 0 {
		return OrderStatusID{}, errorbase.ErrInvalidIdField
	}
	return OrderStatusID{value: orderStatusID}, nil
}

func (o OrderStatusID) GetOrderStatusID() int {
	return o.value
}

type WarehouseID struct {
	value int
}

func NewWarehouseID(warehouseID int) (WarehouseID, error) {
	if warehouseID <= 0 {
		return WarehouseID{}, errorbase.ErrInvalidIdField
	}
	return WarehouseID{value: warehouseID}, nil
}

func (w WarehouseID) GetWarehouseID() int {
	return w.value
}

package models

import "time"

type PurchaseOrder struct {
	ID              int
	Order_number    string
	Order_date      time.Time
	Tracking_code   string
	Buyer_ID        int
	Carrier_ID      int
	Order_Status_ID int
	Warehouse_ID    int
}

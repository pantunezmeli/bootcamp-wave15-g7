package querys

// PURCHASE
var (
	GetReportPurchaseOrders     = "SELECT  b.id, b.id_card_number, b.first_name,  b.last_name, COUNT(ph.id) AS total FROM buyers b LEFT JOIN purchase_orders ph ON b.id = ph.buyer_id GROUP BY b.id, b.first_name, b.last_name ORDER BY total DESC;"
	GetReportPurchaseOrdersById = "SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(po.id) FROM buyers b LEFT JOIN purchase_orders po ON b.id = po.buyer_id WHERE b.id = ?  GROUP BY b.id, b.first_name, b.last_name;"
	CreatePurchaseOrder         = "INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) VALUES (?, ?, ?, ?, ?, ?, ?)"

	CarrierExist      = "SELECT 1 FROM carriers WHERE id = ?;"
	OrderStatusExist  = "SELECT 1 FROM order_status WHERE id = ?;"
	WarehouseExist    = "SELECT 1 FROM warehouses WHERE id = ?;"
	OrderNumberExist  = "SELECT 1 FROM purchase_orders WHERE order_number like ?;"
	TrackingCodeExist = "SELECT 1 FROM purchase_orders WHERE tracking_code like ?;"
	BuyerExist        = "SELECT 1 FROM buyers WHERE id = ?;"
)

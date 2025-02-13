package querys

// BUYERS
var (
	SelectAllBuyers = "SELECT id, id_card_number, first_name, last_name FROM buyers"
	SelectByID      = "SELECT id, id_card_number, first_name, last_name FROM buyers WHERE id = ?"
	InsertBuyer     = "INSERT INTO buyers (id_card_number, first_name, last_name) VALUES (?, ?, ?)"
	DeleteBuyer     = "DELETE FROM buyers WHERE id = ?"
	ExistsBuyer     = "SELECT EXISTS(SELECT 1 FROM buyers WHERE id = ?)"
	ExistCardID     = "SELECT EXISTS(SELECT 1 FROM buyers WHERE id_card_number like ?)"
	SameBuyer       = "SELECT id_card_number FROM buyers WHERE id = ?"
)

// PURCHASE
var (
	GetReportPurchaseOrders     = "SELECT  b.id, b.id_card_number, b.first_name,  b.last_name, COUNT(ph.id) AS total FROM buyers b LEFT JOIN purchase_orders ph ON b.id = ph.buyer_id GROUP BY b.id, b.first_name, b.last_name ORDER BY total DESC;"
	GetReportPurchaseOrdersById = "SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(po.id) FROM buyers b LEFT JOIN purchase_orders po ON b.id = po.buyer_id WHERE b.id = ?  GROUP BY b.id, b.first_name, b.last_name;"
)

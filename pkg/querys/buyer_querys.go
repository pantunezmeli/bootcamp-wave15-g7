package querys

var (
	SelectAllBuyers = "SELECT id, id_card_number, first_name, last_name FROM buyers"
	SelectByID      = "SELECT id, id_card_number, first_name, last_name FROM buyers WHERE id = ?"
	InsertBuyer     = "INSERT INTO buyers (id_card_number, first_name, last_name) VALUES (?, ?, ?)"
	DeleteBuyer     = "DELETE FROM buyers WHERE id = ?"
	ExistsBuyer     = "SELECT EXISTS(SELECT 1 FROM buyers WHERE id = ?)"
	ExistCardID     = "SELECT EXISTS(SELECT 1 FROM buyers WHERE id_card_number like ?)"
)

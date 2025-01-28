package dto

type EmployeeDoc struct {
	Id          int    `json:"id"`
	CardNumber  string `json:"card_number_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WarehouseId int    `json:"warehouse_id"`
}

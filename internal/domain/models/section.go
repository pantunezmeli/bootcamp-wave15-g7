package models

// Section is a struct that represents a Section
type Section struct {
	Id                  int
	Section_number      int
	Current_temperature int
	Minimum_temperature int
	Current_capacity    int
	Minimum_capacity    int
	Maximim_capacity    int
	Warehouse_id        int
	Product_type_id     int
}

// SectionDoc is a struct that represents a Section in JSON format
type SectionDoc struct {
	Id                  int `json:"id"`
	Section_number      int `json:"section_number"`
	Current_temperature int `json:"current_temperature"`
	Minimum_temperature int `json:"minimum_temperature"`
	Current_capacity    int `json:"current_capacity"`
	Minimum_capacity    int `json:"minimum_capacity"`
	Maximim_capacity    int `json:"maximum_capacity"`
	Warehouse_id        int `json:"warehouse_id"`
	Product_type_id     int `json:"product_type_id"`
}

package models

// Section is a struct that represents a Section
type Section struct {
	Id                  int
	Section_Number      int
	Current_Temperature int
	Minimum_Temperature int
	Current_Capacity    int
	Minimum_Capacity    int
	Maximum_Capacity    int
	Warehouse_Id        int
	Product_Type_Id     int
}

// SectionDoc is a struct that represents a Section in JSON format
type SectionDoc struct {
	Id                  int `json:"id"`
	Section_Number      int `json:"section_number"`
	Current_Temperature int `json:"current_temperature"`
	Minimum_Temperature int `json:"minimum_temperature"`
	Current_Capacity    int `json:"current_capacity"`
	Minimum_Capacity    int `json:"minimum_capacity"`
	Maximum_Capacity    int `json:"maximum_capacity"`
	Warehouse_Id        int `json:"warehouse_id"`
	Product_Type_Id     int `json:"product_type_id"`
}

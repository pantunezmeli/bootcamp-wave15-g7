package models

// Section is a struct that represents a Section
type Section struct {
	Id                  int
	Section_Number      string
	Current_Temperature float64
	Minimum_Temperature float64
	Current_Capacity    int
	Minimum_Capacity    int
	Maximum_Capacity    int
	Warehouse_Id        int
	Product_Type_Id     int
}

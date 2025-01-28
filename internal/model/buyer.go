package model

type Buyer struct {
	Id int
	PersonAtributes
}

type BuyerResponse struct {
	Id             int    `json:"id"`
	Card_Number_Id int    `json:"card_number_id"`
	First_Name     string `json:"first_name"`
	Last_Name      string `json:"last_name"`
}

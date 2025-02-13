package productrecords

type Data struct {
	Data ProductRecordsDto
}

type ProductRecordsDto struct {
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductID      int     `json:"product_id"`
}

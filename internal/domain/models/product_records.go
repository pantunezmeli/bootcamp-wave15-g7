package models

import (
	vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	"time"
)

type ProductRecords struct {
	Id             vo.ProductRecordsId
	LastUpdateDate time.Time
	PurchasePrice  float64
	SalePrice      float64
	ProductId      vo.ProductId
}

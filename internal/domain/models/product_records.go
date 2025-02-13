package models

import (
	vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	"time"
)

type ProductRecords struct {
	Id             vo.Id
	LastUpdateDate time.Time
	PurchasePrice  float64
	SalePrice      float64
	ProductID      vo.Id
}

package models

import vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"

type RecordsData struct {
	ProductId    vo.ProductId
	Description  string
	RecordsCount int
}

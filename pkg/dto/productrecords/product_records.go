package productrecords

import (
	m "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product/errsv"
	"time"
)

const LAYOUT_DATE = "2006-01-02"

type Data struct {
	Data ProductRecordsDto
}

type ProductRecordsDto struct {
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductID      int     `json:"product_id"`
}

func ValidAndParserDTO(dto ProductRecordsDto, record *m.ProductRecords) error {
	id, err := vo.NewId(dto.ProductID)
	if err != nil {
		return errsv.ErrValidEntity{Message: "Invalid product id"}
	}

	if dto.LastUpdateDate == "" {
		return errsv.ErrValidEntity{Message: "Invalid last update date"}
	}

	if dto.PurchasePrice <= 0 || dto.SalePrice <= 0 {
		return errsv.ErrValidEntity{Message: "Invalid price"}
	}

	date, err := time.Parse(LAYOUT_DATE, dto.LastUpdateDate)
	if err != nil {
		return errsv.ErrValidEntity{Message: "Invalid date"}
	}

	record.LastUpdateDate = date
	record.PurchasePrice = dto.PurchasePrice
	record.SalePrice = dto.SalePrice
	record.ProductId = id

	return nil
}

func ParserRecordsToDto(record m.ProductRecords) ProductRecordsDto {
	return ProductRecordsDto{
		LastUpdateDate: record.LastUpdateDate.Format("2006-01-02"),
		PurchasePrice:  record.PurchasePrice,
		SalePrice:      record.SalePrice,
		ProductID:      record.ProductId.GetId(),
	}
}

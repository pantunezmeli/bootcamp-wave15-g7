package productrecords

import m "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type RecordsResponse struct {
	ProductId    int    `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int    `json:"records_count"`
}

func ParserRecordsDataToDto(records []m.RecordsData) (recordsResponse []RecordsResponse) {
	for _, record := range records {
		recordsResponse = append(recordsResponse, RecordsResponse{
			ProductId:    record.ProductId.GetId(),
			Description:  record.Description,
			RecordsCount: record.RecordsCount,
		})
	}
	return
}

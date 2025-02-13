package product_records

import (
	"errors"
	rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/productrecords"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product/errsv"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/productrecords"
)

type ProductRecordsService struct {
	rp rp.IProductRecordsRepository
}

func NewProductRecordsService(rp rp.IProductRecordsRepository) *ProductRecordsService {
	return &ProductRecordsService{rp: rp}
}

func (p ProductRecordsService) CreateProductRecord(newProductRecord dto.ProductRecordsDto) (dto.ProductRecordsDto, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRecordsService) GetProductRecord(productID int) (recordsDto []dto.RecordsResponse, err error) {

	recordsData, errGet := p.rp.GetRecordsDataOptionalId(&productID)
	if errGet != nil {
		if errors.Is(errGet, rp.ErrRecordsNotFound) {
			err = errsv.ErrNotFoundProduct{Message: "Records not found"}
			return
		}
		err = errsv.ErrProduct{Message: "Error searching Records Product"}
		return
	}

	recordsDto = dto.ParserRecordsDataToDto(recordsData)
	return
}

package product_records

import (
	"errors"
	m "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	errdb "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product/errordb"
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

func (p ProductRecordsService) CreateProductRecord(newProductRecord dto.ProductRecordsDto) (recordDto dto.ProductRecordsDto, err error) {
	var newRecord m.ProductRecords
	errValid := dto.ValidAndParserDTO(newProductRecord, &newRecord)
	if errValid != nil {
		err = errsv.ErrValidEntity{Message: errValid.Error()}
		return
	}

	if errCreate := p.rp.CreateProductRecord(&newRecord); errCreate != nil {
		if errors.As(errCreate, &errdb.ErrViolateFK{}) {
			err = errsv.ErrConflict{Message: errCreate.Error()}
			return
		}
		err = errsv.ErrService{Message: "Error creating Product Record"}
		return
	}

	recordDto = dto.ParserRecordsToDto(newRecord)
	return
}

func (p ProductRecordsService) GetProductRecord(productID *int) (recordsDto []dto.RecordsResponse, err error) {

	recordsData, errGet := p.rp.GetRecordsDataOptionalId(productID)
	if errGet != nil {
		if errors.Is(errGet, rp.ErrRecordsNotFound) {
			err = errsv.ErrNotFoundEntity{Message: "Records not found"}
			return
		}
		err = errsv.ErrService{Message: "Error searching Records Product"}
		return
	}

	recordsDto = dto.ParserRecordsDataToDto(recordsData)
	return
}

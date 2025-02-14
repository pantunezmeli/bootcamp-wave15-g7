package carrier

import (
	"errors"
	"fmt"

	customErrors "github.com/pantunezmeli/bootcamp-wave15-g7/internal/errors"
	repository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/carrier"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/carrier"
)

var (
	ErrLocalityNotExists = errors.New("error locality does not exists")
)

type CarrierService struct {
	rp repository.ICarrierRepository
}

func NewCarrierService(rp repository.ICarrierRepository) *CarrierService {
	return &CarrierService{rp: rp}
}

// ! 1)
func (s *CarrierService) AddCarrier(req dto.CarrierDoc) (c dto.CarrierDoc, err error) {

	carrierModel, err := req.ConvertToModel(req)
	if err != nil {
		return dto.CarrierDoc{}, customErrors.ErrInvalidParameter{Parameter: err.Error()}
	}

	createdCarrier, err := s.rp.AddCarrierToDB(carrierModel)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrForeignKeyViolation):
			return dto.CarrierDoc{}, customErrors.ErrForeignKey{Err: err}
		case errors.Is(err, customErrors.ErrDuplicateEntry):
			return dto.CarrierDoc{}, customErrors.ErrDuplicate{Err: err}
		case errors.Is(err, customErrors.ErrDBGenericError):
			return dto.CarrierDoc{}, customErrors.ErrDatabase{Err: err}
		case errors.Is(err, customErrors.ErrInsertingData):
			return dto.CarrierDoc{}, customErrors.ErrDatabase{Err: err}
		case errors.Is(err, customErrors.ErrGettingLastID), errors.Is(err, customErrors.ErrConvertingID):
			return dto.CarrierDoc{}, customErrors.ErrDatabase{Err: err}
		default:
			return dto.CarrierDoc{}, customErrors.ErrDatabase{Err: fmt.Errorf("unexpected error: %w", err)}
		}
	}

	carrierDTO, err := dto.CarrierDoc{}.ConvertToDTO(createdCarrier)
	if err != nil {
		return dto.CarrierDoc{}, customErrors.ErrConvertion{Err: err}
	}

	return carrierDTO, nil
}

// ! 2)
func (s *CarrierService) GetCarriesAmount(id *int) (result []dto.CarrierByLocalityID, err error) {
	carriesAmount, err := s.rp.GetCarriesAmountByLocalityID(id)
	if err != nil {
		if errors.Is(err, customErrors.ErrLocalityNotFound) {
			return []dto.CarrierByLocalityID{}, customErrors.ErrNotFound{Err: err}
		}
		return []dto.CarrierByLocalityID{}, customErrors.ErrDatabase{Err: fmt.Errorf("unexpected error: %w", err)}
	}

	var carriesAmountsDTOs []dto.CarrierByLocalityID
	for _, item := range carriesAmount {
		dtoItem, err := dto.CarrierByLocalityID{}.ConvertToDTO(item)
		if err != nil {
			return nil, customErrors.ErrConvertion{Err: err}
		}
		carriesAmountsDTOs = append(carriesAmountsDTOs, dtoItem)
	}

	return carriesAmountsDTOs, nil
}

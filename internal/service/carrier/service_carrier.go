package carrier

import (
	"errors"
	"fmt"

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
		return dto.CarrierDoc{}, ErrInvalidParameter{Parameter: err.Error()}
	}

	createdCarrier, err := s.rp.AddCarrierToDB(carrierModel)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrForeignKeyViolation):
			return dto.CarrierDoc{}, ErrForeignKey{Err: err}
		case errors.Is(err, repository.ErrDuplicateEntry):
			return dto.CarrierDoc{}, ErrDuplicate{Err: err}
		case errors.Is(err, repository.ErrDBGenericError):
			return dto.CarrierDoc{}, ErrDatabase{Err: err}
		case errors.Is(err, repository.ErrInsertingData):
			return dto.CarrierDoc{}, ErrDatabase{Err: err}
		case errors.Is(err, repository.ErrGettingLastID), errors.Is(err, repository.ErrConvertingID):
			return dto.CarrierDoc{}, ErrDatabase{Err: err}
		default:
			return dto.CarrierDoc{}, ErrDatabase{Err: fmt.Errorf("unexpected error: %w", err)}
		}
	}

	carrierDTO, err := dto.CarrierDoc{}.ConvertToDTO(createdCarrier)
	if err != nil {
		return dto.CarrierDoc{}, ErrConvertion{Err: err}
	}

	return carrierDTO, nil
}

// ! 2)
func (s *CarrierService) GetCarriesAmount(id int) (result dto.CarrierByLocalityID, err error) {
	carriesAmount, err := s.rp.GetCarriesAmountByLocalityID(id)
	if err != nil {
		if errors.Is(err, repository.ErrLocalityNotFound) {
			return dto.CarrierByLocalityID{}, ErrLocalityNotExists
		}
		return dto.CarrierByLocalityID{}, ErrDatabase{Err: fmt.Errorf("unexpected error: %w", err)}
	}

	carriesAmountDTO, err := dto.CarrierByLocalityID{}.ConvertToDTO(carriesAmount)
	if err != nil {
		return dto.CarrierByLocalityID{}, ErrConvertion{Err: err}
	}

	return carriesAmountDTO, nil
}

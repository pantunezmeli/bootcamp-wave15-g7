package carrier

import (
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/carrier"
)

type ICarrierService interface {

	// ! 1)
	AddCarrier(req dto.CarrierDoc) (c dto.CarrierDoc, err error)

	// ! 2)
	GetCarriesAmount(id *int) (result []dto.CarrierByLocalityID, err error)
}

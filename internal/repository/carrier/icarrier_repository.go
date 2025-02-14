package carrier

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type ICarrierRepository interface {

	//! 1)
	AddCarrierToDB(carrier models.Carrier) (c models.Carrier, err error)

	//! 2)
	GetCarriesAmountByLocalityID(id *int) ([]CarrierByLocality, error)
}

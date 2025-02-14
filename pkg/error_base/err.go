package errorbase

import "errors"

var (
	ErrInvalidId              = errors.New("invalid id")
	ErrEmptyParameters        = errors.New("empty parameters")
	ErrEmptyList              = errors.New("empty list")
	ErrConflict               = errors.New("element already exist")
	ErrNotFound               = errors.New("not found")
	ErrModelInvalid           = errors.New("model invalid")
	ErrStorageOperationFailed = errors.New("storage operation failed")
	ErrInvalidNumber          = errors.New("invalid number")
	ErrInvalidRequest         = errors.New("invalid request")
	ErrUnprocessable          = errors.New("incorrect request")

	// Product Batch Errors
	ErrInvalidBatchNumber        = errors.New("invalid batch number")
	ErrInvalidCurrentQuantity    = errors.New("invalid current quantity")
	ErrInvalidCurrentTemperature = errors.New("invalid current temperature")
	ErrInvalidDueDate            = errors.New("invalid due date")
	ErrInvalidInitialQuantity    = errors.New("invalid initial quantity")
	ErrInvalidManufacturingDate  = errors.New("invalid manufacturing date")
	ErrInvalidManufacturingHour  = errors.New("invalid manufacturing hour")
	ErrInvalidMinumumTemperature = errors.New("invalid minimum temperature")
	ErrInvalidProductID          = errors.New("invalid product id")
	ErrInvalidSectionID          = errors.New("invalid section id")
	ErrInvalidProductBatch       = errors.New("invalid product batch")
	ErrInvalidProductBatchID     = errors.New("invalid product batch id")
	ErrInvalidProductBatchNumber = errors.New("invalid product batch number")
	ErrInvalidProductBatchList   = errors.New("invalid product batch list")

	// Section Errors
	ErrInvalidSectionNumber         = errors.New("invalid section number")
	ErrInvalidMinimumTemperature    = errors.New("invalid minimum temperature")
	ErrInvalidCurrentCapacity       = errors.New("invalid current capacity")
	ErrInvalidMinimumCapacity       = errors.New("invalid minimum capacity")
	ErrInvalidMaximumCapacity       = errors.New("invalid maximum capacity")
	ErrInvalidWarehouseID           = errors.New("invalid warehouse id")
	ErrInvalidProductTypeID         = errors.New("invalid product type id")
	ErrInvalidSection               = errors.New("invalid section")
	ErrInvalidSectionList           = errors.New("invalid section list")
	ErrInvalidSectionCurrentTemp    = errors.New("invalid section current temperature")
	ErrInvalidSectionMinTemp        = errors.New("invalid section minimum temperature")
	ErrInvalidSectionCurrentCap     = errors.New("invalid section current capacity")
	ErrInvalidSectionMinCap         = errors.New("invalid section minimum capacity")
	ErrInvalidSectionMaxCap         = errors.New("invalid section maximum capacity")
	ErrInvalidSectionWarehouseID    = errors.New("invalid section warehouse id")
	ErrInvalidSectionProductTypeID  = errors.New("invalid section product type id")
	ErrInvalidSectionProductBatch   = errors.New("invalid section product batch")
	ErrInvalidSectionProductBatchID = errors.New("invalid section product batch id")
)

type ErrInvalidParameter struct {
	Parameter string
}

func (e ErrInvalidParameter) Error() string {

	return "invalid parameter: " + e.Parameter

}

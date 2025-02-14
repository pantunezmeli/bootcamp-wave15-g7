package value_objects

import "errors"

var (
	ErrWarehouseInvalidId         = errors.New("invalid id")
	ErrInvalidWareHouseCode       = errors.New("invalid warehouse code")
	ErrInvalidWarehouseTelephone  = errors.New("invalid telephone")
	ErrInvalidWarehouseAddress    = errors.New("invalid address")
	ErrInvalidLocalityWarehouseID = errors.New("invalid locality id")
)

// * ##################### Id ######################

type IDWarehouse int64

func NewWarehouseId(value int) (id IDWarehouse, err error) {
	if value <= 0 {
		return 0, ErrWarehouseInvalidId
	}
	return IDWarehouse(value), nil
}

// * ##################### WareHouse_Code ######################

type WareHouseCode string

func NewWareHouseCode(value string) (code WareHouseCode, err error) {
	if len(value) < 3 {
		return "", ErrInvalidWareHouseCode
	}
	return WareHouseCode(value), nil
}

// * ##################### Address ######################

type WarehouseAddress string

func NewWarehouseAddress(value string) (address WarehouseAddress, err error) {
	if len(value) < 5 {
		return "", ErrInvalidWarehouseAddress
	}
	return WarehouseAddress(value), nil
}

// * ##################### Telephone ######################

type WarehouseTelephone string

func NewWarehouseTelephone(value string) (phone WarehouseTelephone, err error) {
	if len(value) < 10 {
		return "", ErrInvalidWarehouseTelephone
	}
	return WarehouseTelephone(value), nil
}

// * ##################### LocalityID ######################

type LocalityID int

func NewWarehouseLocalityID(value int) (locality LocalityID, err error) {
	if value < 0 {
		return 0, ErrInvalidLocalityWarehouseID
	}
	return LocalityID(value), nil
}

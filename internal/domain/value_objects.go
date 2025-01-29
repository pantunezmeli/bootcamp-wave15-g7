package domain

import "errors"

var (
	ErrInvalidId              = errors.New("invalid id")
	ErrInvalidWareHouseCode   = errors.New("invalid warehouse code")
	ErrInvalidTelephone       = errors.New("invalid telephone")
	ErrInvalidMinimunCapacity = errors.New("invalid minimun capacity")
	ErrMinimunTemperature     = errors.New("invalid minimun temperature")
	ErrInvalidAddress         = errors.New("invalid address")
)

// * ##################### Id ######################
// Estructura
type Id struct {
	value int
}

// Validación
func NewId(value int) (id Id, err error) {
	if value <= 0 {
		return Id{}, ErrInvalidId
	}
	return Id{value: value}, nil
}

// Obtener el valor
func (id Id) GetId() int {
	return id.value
}

// * ##################### WareHouse_Code ######################

type WareHouseCode struct {
	value string
}

func NewWareHouseCode(value string) (code WareHouseCode, err error) {
	if len(value) < 3 {
		return WareHouseCode{}, ErrInvalidWareHouseCode
	}
	return WareHouseCode{value: value}, nil
}

func (code WareHouseCode) GetWareHouseCode() string {
	return code.value
}

// * ##################### Address ######################

type Address struct {
	value string
}

func NewAddress(value string) (address Address, err error) {
	if len(value) < 5 {
		return Address{}, ErrInvalidAddress
	}
	return Address{value: value}, nil
}

func (address Address) GetAddress() string {
	return address.value
}

// * ##################### Telephone ######################

type Telephone struct {
	value string
}

func NewTelephone(value string) (phone Telephone, err error) {
	if len(value) < 10 {
		return Telephone{}, ErrInvalidTelephone
	}
	return Telephone{value: value}, nil
}

func (phone Telephone) GetTelephone() string {
	return phone.value
}

// * ##################### minimun_capacity ######################

type MinimunCapacity struct {
	value int
}

func NewMinimunCapacity(value int) (minCapacity MinimunCapacity, err error) {
	if value <= 0 {
		return MinimunCapacity{}, ErrInvalidMinimunCapacity
	}
	return MinimunCapacity{value: value}, nil
}

func (minCapacity MinimunCapacity) GetMinimunCapacity() int {
	return minCapacity.value
}

// * ##################### minimun_temperature ######################

type MinimunTemperature struct {
	value int
}

func NewMinimunTemperature(value int) (minTemperature MinimunTemperature, err error) {
	if value <= -100 {
		return MinimunTemperature{}, ErrMinimunTemperature
	}
	return MinimunTemperature{value: value}, nil
}

func (minTemperature MinimunTemperature) GetMinimunTemperature() int {
	return minTemperature.value
}

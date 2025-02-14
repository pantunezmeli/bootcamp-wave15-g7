package value_objects

import "errors"

var (
	ErrInvalidId = errors.New("invalid id")
	//ErrInvalidWareHouseCode   = errors.New("invalid warehouse code")
	ErrInvalidTelephone       = errors.New("invalid telephone")
	ErrInvalidMinimunCapacity = errors.New("invalid minimun capacity")
	ErrMinimunTemperature     = errors.New("invalid minimun temperature")
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidName            = errors.New("invalid name")
	ErrInvalidCardNumber      = errors.New("invalid card number")
)

// * ##################### Id ######################
// Estructura

type ProductId int
type ProductRecordsId int
type ProductTypeId int

func NewProductTypeId(value int) (id ProductTypeId, err error) {
	if value <= 0 {
		return 0, ErrInvalidId
	}
	return ProductTypeId(value), nil
}

func NewProductId(value int) (id ProductId, err error) {
	if value <= 0 {
		return 0, ErrInvalidId
	}
	return ProductId(value), nil
}

func NewProductRecordsId(value int) (id ProductRecordsId, err error) {
	if value <= 0 {
		return 0, ErrInvalidId
	}
	return ProductRecordsId(value), nil
}

type Id struct {
	value int
}

type CardNumber struct {
	value string
}
type Name struct {
	value string
}

// Validación
func NewId(value int) (id Id, err error) {
	if value <= 0 {
		return Id{}, ErrInvalidId
	}
	return Id{value: value}, nil
}
func NewCardNumber(value string) (cardNumber CardNumber, err error) {
	if value == "" {
		return CardNumber{}, ErrInvalidCardNumber
	}
	return CardNumber{value: value}, nil
}
func NewName(value string) (name Name, err error) {
	if value == "" {
		return Name{}, ErrInvalidName
	}
	return Name{value: value}, nil
}

// Seteo sin validación
func NewOptionalId(value int) Id {
	return Id{value: value}
}
func NewOptionalCardNumber(value string) CardNumber {
	return CardNumber{value: value}
}
func NewOptionalName(value string) Name {
	return Name{value: value}
}

// Obtener el valor
func (id Id) GetId() int {
	return id.value
}
func (cardNumber CardNumber) GetCardNumber() string {
	return cardNumber.value
}
func (name Name) GetName() string {
	return name.value
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

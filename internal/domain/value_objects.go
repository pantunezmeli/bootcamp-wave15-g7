package domain

import "errors"

var ErrInvalidId = errors.New("invalid id")
var ErrInvalidCardNumber = errors.New("invalid card number")
var ErrInvalidName = errors.New("invalid name")

// Estructura
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
	if value < 0 {
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

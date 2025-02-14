package country

import "errors"

var (
	ErrInvalidCountryId = errors.New("invalid Country id")
	ErrInvalidCountryName = errors.New("invalid Country name")
)



type CountryId int

func NewCountryId(value int) (Country CountryId, err error) {
	if value < 0 {
		err = ErrInvalidCountryId
		return
	}
	Country = CountryId(value)
	return
}


type CountryName string

func NewCountryName(value string) (Country CountryName, err error) {
	if len(value) < 2 || len(value) > 255 {
		err = ErrInvalidCountryName
		return
	}
	Country = CountryName(value)
	return
}

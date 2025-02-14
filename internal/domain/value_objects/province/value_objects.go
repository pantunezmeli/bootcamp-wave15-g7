package province

import "errors"

var (
	ErrInvalidProvinceId = errors.New("invalid province id")
	ErrInvalidProvinceName = errors.New("invalid province name")
)

type ProvinceId int

func NewProvinceId(value int) (Province ProvinceId, err error) {
	if value < 0 {
		err = ErrInvalidProvinceId
		return
	}
	Province = ProvinceId(value)
	return
}


type ProvinceName string

func NewProvinceName(value string) (Province ProvinceName, err error) {
	if len(value) < 2 || len(value) > 255 {
		err = ErrInvalidProvinceName
		return
	}
	Province = ProvinceName(value)
	return
}

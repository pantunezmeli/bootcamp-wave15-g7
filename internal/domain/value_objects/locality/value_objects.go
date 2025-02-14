package locality

import "errors"

var (
	ErrInvalidLocalityId = errors.New("invalid locality id")
	ErrInvalidLocalityName = errors.New("invalid locality name")
)



type LocalityId int

func NewLocalityId(value int) (locality LocalityId, err error) {
	if value < 0 {
		err = ErrInvalidLocalityId
		return
	}
	locality = LocalityId(value)
	return
}


type LocalityName string

func NewLocalityName(value string) (locality LocalityName, err error) {
	if len(value) < 2 || len(value) > 255 {
		err = ErrInvalidLocalityName
		return
	}
	locality = LocalityName(value)
	return
}

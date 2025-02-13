package locality

import (
	"fmt"

	locality_dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/locality"
)

var (
	NameString = "locality_name"
	ProvinceIdString = "province_id"
)


type ErrMissingParameters struct {
	missingParameter string;
}

func(e *ErrMissingParameters) Error() string {
	return fmt.Sprintf("%s is needed", e.missingParameter)
}

type ErrInvalidParameter struct {
	message string;
}

func(e *ErrInvalidParameter) Error() string {
	return e.message
}

type LocalityService interface {
	Save(locality_dto.LocalityRequest) (localityDto locality_dto.LocalityDoc, err error)
	GetById(id int) (localityDto locality_dto.LocalityDoc, err error)
}


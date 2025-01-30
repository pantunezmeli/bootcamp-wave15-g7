package seller

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
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

type SellerService interface {
	GetAll() (sellers []dto.SellerDoc, err error)
	GetById(int) (seller dto.SellerDoc, err error)
	Save(dto.SellerDoc) (seller dto.SellerDoc, err error)
	Delete(id int) (err error)
	Update(reqBody dto.SellerDoc) (seller dto.SellerDoc, err error)
}

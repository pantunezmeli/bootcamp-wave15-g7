package seller

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)


var (
	ErrMissingParameters = errors.New("missing parameters")
)

type SellerService interface {
	GetAll() (sellers []dto.SellerDoc, err error)
	GetById(int) (seller dto.SellerDoc, err error)
	Save(dto.SellerDoc) (seller dto.SellerDoc, err error)
	Delete(id int) (err error)
}

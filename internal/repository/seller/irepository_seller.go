package seller

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

var (
	ErrSellerNotFound = errors.New("seller not found")
)


type SellerRepository interface {
	GetAll() (sellers []models.Seller, err error)
	GetById(id int) (seller models.Seller, err error)
}
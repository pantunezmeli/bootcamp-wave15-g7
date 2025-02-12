package seller

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

var (
	ErrSellerNotFound = errors.New("seller not found")
	ErrCidAlreadyExists = errors.New("cid already exists")
	ErrConnectionError = errors.New("connection error")
)


type SellerRepository interface {
	GetAll() (sellers []models.Seller, err error)
	GetById(int) (seller models.Seller, err error)
	Save(models.Seller) (seller models.Seller, err error)
	Delete(id int) (err error)
	Update(sellerModel models.Seller) (sellerUpdated models.Seller, err error)
}
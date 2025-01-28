package seller

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"




type SellerRepository interface {
	GetAll() (sellers []models.Seller, err error)
}
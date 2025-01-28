package seller

import "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"


type SellerService interface {
	GetAll() (sellers []dto.SellerDoc, err error)
}
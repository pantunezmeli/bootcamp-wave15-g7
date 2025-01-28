package loader

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type SellerLoader interface {
	Load() (v map[int]models.Seller, err error)
}

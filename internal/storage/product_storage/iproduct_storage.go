package productstorage

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IProductLoader interface {
	GetDb() (map[int]models.Product, error)
	RemoveProduct(productID int) error
	SaveProduct(newProduct models.Product) error
}

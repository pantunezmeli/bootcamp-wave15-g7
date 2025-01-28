package product

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"

type IProductService interface {
	GetAll() (map[int]model.Product, error)
}

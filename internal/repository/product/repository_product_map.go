package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
)

func NewProductRepositoryMap(db map[int]model.Product) *ProductRepositoryMap {
	defaultDb := make(map[int]model.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductRepositoryMap{db: defaultDb}
}

type ProductRepositoryMap struct {
	db map[int]model.Product
}

func (p ProductRepositoryMap) GetAll() map[int]model.Product {
	return p.db
}

func (p ProductRepositoryMap) GetByID(id int) (model.Product, error) {
	product, ok := p.db[id]
	if !ok {
		return model.Product{}, ErrProductNotFound
	}
	return product, nil
}

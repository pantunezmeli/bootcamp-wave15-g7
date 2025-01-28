package product

import (
	"errors"
	"fmt"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader/product"
)

func NewProductRepositoryMap(loader product.IProductLoader) *ProductRepositoryMap {
	return &ProductRepositoryMap{loader: loader}
}

type ProductRepositoryMap struct {
	loader product.IProductLoader
}

func (p ProductRepositoryMap) GetAll() (map[int]model.Product, error) {
	products, err := p.loader.GetDb()
	if err != nil {
		fmt.Println(err)
	}
	return products, err
}

func (p ProductRepositoryMap) GetByID(id int) (model.Product, error) {
	products, err := p.GetAll()
	if err != nil {
		return model.Product{}, errors.New("Error getting all products")
	}

	productSearch, ok := products[id]
	if !ok {
		return model.Product{}, ErrProductNotFound
	}

	return productSearch, nil
}

package product

import (
	"errors"
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/product_storage"
)

func NewProductRepositoryMap(storage productstorage.IProductLoader) *ProductRepositoryMap {
	return &ProductRepositoryMap{storage: storage}
}

type ProductRepositoryMap struct {
	storage productstorage.IProductLoader
}

func (p ProductRepositoryMap) GetAll() (map[int]models.Product, error) {
	products, err := p.storage.GetDb()
	if err != nil {
		fmt.Println(err)
	}
	return products, err
}

func (p ProductRepositoryMap) GetByID(id int) (models.Product, error) {
	products, err := p.GetAll()
	if err != nil {
		return models.Product{}, errors.New("Error getting all products")
	}

	productSearch, ok := products[id]
	if !ok {
		return models.Product{}, ErrProductNotFound
	}

	return productSearch, nil
}

func (p ProductRepositoryMap) DeleteProduct(id int) error {
	productSearch, errSearch := p.GetByID(id)
	if errSearch != nil {
		return errSearch
	}

	errDelete := p.storage.RemoveProduct(productSearch.ID.GetId())
	if errDelete != nil {
		return ErrProductRepository{msg: "Error deleting product"}
	}

	return nil
}

func (p ProductRepositoryMap) CreateProduct(product models.Product) error {
	return p.storage.SaveProduct(product)
}

func (p ProductRepositoryMap) ProductCodeExist(productCode string) bool {
	products, _ := p.GetAll()

	for _, productMap := range products {
		if productMap.ProductCode == productCode {
			return true
		}
	}
	return false
}

func (p ProductRepositoryMap) GetLastID() value_objects.Id {
	products, _ := p.GetAll()
	var lastId int
	for _, productMap := range products {
		if lastId < productMap.ID.GetId() {
			lastId = productMap.ID.GetId()
		}
	}
	id, _ := value_objects.NewId(lastId + 1)
	return id
}

func (p ProductRepositoryMap) UpdateProduct(product models.Product) error {
	return p.storage.SaveProduct(product)
}

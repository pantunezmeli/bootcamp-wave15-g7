package product

import (
	"encoding/json"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	"os"
)

type IProductLoader interface {
	GetDb() (map[int]model.Product, error)
	RemoveProduct(productID int) error
	AddProduct(newProduct model.Product) error
}

type ProductJSONFile struct {
	path string
}

func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{path: path}
}

func (l *ProductJSONFile) GetDb() (map[int]model.Product, error) {
	file, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ProductJson []dto.ProductDTO
	err = json.NewDecoder(file).Decode(&ProductJson)
	if err != nil {
		return nil, err
	}

	list := make(map[int]model.Product)
	for _, v := range ProductJson {
		list[v.ID] = dto.ParseDTOProduct(v)
	}

	return list, nil
}

func (l *ProductJSONFile) AddProduct(newProduct model.Product) error {
	products, err := l.GetDb()
	if err != nil {
		return err
	}

	products[newProduct.ID] = newProduct

	return l.saveToSlice(products)
}

func (l *ProductJSONFile) RemoveProduct(productID int) error {
	products, err := l.GetDb()
	if err != nil {
		return err
	}

	delete(products, productID)

	return l.saveToSlice(products)
}

func (l *ProductJSONFile) saveToSlice(products map[int]model.Product) error {
	var productSlice []dto.ProductDTO
	for _, v := range products {
		productSlice = append(productSlice, dto.ParserProductToDTO(v))
	}

	return l.save(productSlice)
}

func (l *ProductJSONFile) save(products []dto.ProductDTO) error {
	file, err := os.Create(l.path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(products)
}

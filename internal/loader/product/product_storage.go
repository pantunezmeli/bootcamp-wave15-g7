package product

import (
	"encoding/json"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product"
	"os"
)

type IProductLoader interface {
	GetDb() (map[int]models.Product, error)
	RemoveProduct(productID int) error
	SaveProduct(newProduct models.Product) error
}

type ProductJSONFile struct {
	path string
}

func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{path: path}
}

func (l *ProductJSONFile) GetDb() (map[int]models.Product, error) {
	file, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ProductJson []product.ProductDTO
	err = json.NewDecoder(file).Decode(&ProductJson)
	if err != nil {
		return nil, err
	}

	list := make(map[int]models.Product)
	for _, v := range ProductJson {
		var newProduct models.Product
		product.ValidAndParserDTO(v, &newProduct)
		list[v.ID] = newProduct
	}

	return list, nil
}

func (l *ProductJSONFile) SaveProduct(newProduct models.Product) error {
	products, err := l.GetDb()
	if err != nil {
		return err
	}

	products[newProduct.ID.GetId()] = newProduct

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

func (l *ProductJSONFile) saveToSlice(products map[int]models.Product) error {
	var productSlice []product.ProductDTO
	for _, v := range products {
		productSlice = append(productSlice, product.ParserProductToDTO(v))
	}

	return l.save(productSlice)
}

func (l *ProductJSONFile) save(products []product.ProductDTO) error {
	file, err := os.Create(l.path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(products)
}

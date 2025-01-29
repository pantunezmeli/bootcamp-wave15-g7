package product

import (
	"errors"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	product2 "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product"
)

func NewProductService(rp product.IProductRepository) *ProductService {
	return &ProductService{rp: rp}
}

type ProductService struct {
	rp product.IProductRepository
}

func (p ProductService) UpdateProduct(id int, patch product2.UpdateProductRequest) (productUpdate product2.ProductDTO, err error) {
	productSearch, errSearch := p.rp.GetByID(id)
	if errSearch != nil {
		if errors.Is(errSearch, product.ErrProductNotFound) {
			err = ErrNotFoundProduct{message: "Product not found"}
			return
		}
		err = ErrProduct{message: "Error searching product"}
		return
	}

	errPatch := product2.PatchProduct(patch, &productSearch)
	if errPatch != nil {
		err = ErrValidProduct{message: errPatch.Error()}
		return
	}

	errSave := p.rp.UpdateProduct(productSearch)
	if errSave != nil {
		err = ErrProduct{message: "Error save product"}
		return
	}

	return product2.ParserProductToDTO(productSearch), nil
}

func (p ProductService) GetAll() ([]product2.ProductDTO, error) {
	products, errSearch := p.rp.GetAll()
	if errSearch != nil {
		return make([]product2.ProductDTO, 0), errSearch
	}
	return product2.ParserMapProductToListDTO(products), nil
}

func (p ProductService) GetByID(id int) (productDTO product2.ProductDTO, err error) {
	productSearch, errSearch := p.rp.GetByID(id)
	if errSearch != nil {
		if errors.Is(errSearch, product.ErrProductNotFound) {
			err = ErrNotFoundProduct{message: "Product not found"}
			return
		}
		err = ErrProduct{message: "Error searching product"}
		return
	}

	productDTO = product2.ParserProductToDTO(productSearch)
	return
}

func (p ProductService) CreateProduct(product product2.ProductDTO) (productDto product2.ProductDTO, err error) {

	var newProduct models.Product
	errValid := product2.ValidAndParserDTO(product, &newProduct)
	if errValid != nil {
		err = ErrValidProduct{message: errValid.Error()}
		return
	}

	if p.rp.ProductCodeExist(newProduct.ProductCode) {
		err = ErrValidProduct{message: "Product code already exists"}
		return
	}

	newProduct.ID = p.rp.GetLastID()

	if errCreate := p.rp.CreateProduct(newProduct); errCreate != nil {
		err = ErrProduct{message: "Error creating product"}
		return
	}

	productDto = product2.ParserProductToDTO(newProduct)
	return
}

func (p ProductService) DeleteProduct(id int) (err error) {
	errDelete := p.rp.DeleteProduct(id)
	if errDelete != nil {
		if errors.Is(errDelete, product.ErrProductNotFound) {
			err = ErrNotFoundProduct{message: "Product not found"}
			return
		}
		err = ErrProduct{message: "Error deleted product"}
	}

	return
}

package product

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

func NewProductService(rp product.IProductRepository) *ProductService {
	return &ProductService{rp: rp}
}

type ProductService struct {
	rp product.IProductRepository
}

func (p ProductService) GetAll() ([]dto.ProductDTO, error) {
	products, errSearch := p.rp.GetAll()
	if errSearch != nil {
		return make([]dto.ProductDTO, 0), errSearch
	}
	return dto.ParserListProductToDTO(products), nil
}

func (p ProductService) GetByID(id int) (dto.ProductDTO, error) {
	productSearch, errSearch := p.rp.GetByID(id)
	if errSearch != nil {
		return dto.ProductDTO{}, errSearch
	}

	return dto.ParserProductToDTO(productSearch), nil
}

func (p ProductService) CreateProduct(product dto.ProductDTO) (productDto dto.ProductDTO, err error) {

	var newProduct model.Product
	errValid := dto.ValidAndParserDTO(product, &newProduct)
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

	productDto = dto.ParserProductToDTO(newProduct)
	return
}

func (p ProductService) DeleteProduct(id int) error {
	return p.rp.DeleteProduct(id)
}

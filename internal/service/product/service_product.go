package product

import (
	"errors"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	errdb "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product/errordb"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product/errsv"
	product2 "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product"
)

func NewProductService(rp rp.IProductRepository) *ProductService {
	return &ProductService{rp: rp}
}

type ProductService struct {
	rp rp.IProductRepository
}

func (p ProductService) UpdateProduct(id int, patch product2.UpdateProductRequest) (productUpdate product2.ProductDTO, err error) {
	productSearch, errSearch := p.rp.GetByID(id)
	if errSearch != nil {
		if errors.Is(errSearch, rp.ErrProductNotFound) {
			err = errsv.ErrNotFoundProduct{Message: "Product not found"}
			return
		}
		err = errsv.ErrProduct{Message: "Error searching product"}
		return
	}

	if patch.ProductCode != nil {
		errValidCode := p.rp.ProductCodeExist(*patch.ProductCode)
		if errValidCode != nil {
			if errors.Is(errValidCode, rp.ErrProductCodeAlreadyExist) {
				err = errsv.ErrProductConflict{Message: "Product code already exists"}
				return
			}
			err = errsv.ErrProduct{Message: "Error valid product code"}
			return
		}
	}

	errPatch := product2.PatchProduct(patch, &productSearch)
	if errPatch != nil {
		err = errsv.ErrValidProduct{Message: errPatch.Error()}
		return
	}

	errUpdate := p.rp.UpdateProduct(productSearch)
	if errUpdate != nil {
		if errors.As(errUpdate, &errdb.ErrViolateFK{}) {
			err = errsv.ErrProductConflict{Message: errUpdate.Error()}
			return
		}

		err = errsv.ErrProduct{Message: "Error save product"}
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
		if errors.Is(errSearch, rp.ErrProductNotFound) {
			err = errsv.ErrNotFoundProduct{Message: "Product not found"}
			return
		}
		err = errsv.ErrProduct{Message: "Error searching product"}
		return
	}

	productDTO = product2.ParserProductToDTO(productSearch)
	return
}

func (p ProductService) CreateProduct(product product2.ProductDTO) (productDto product2.ProductDTO, err error) {

	var newProduct models.Product
	errValid := product2.ValidAndParserDTO(product, &newProduct)
	if errValid != nil {
		err = errsv.ErrValidProduct{Message: errValid.Error()}
		return
	}

	errValidCode := p.rp.ProductCodeExist(newProduct.ProductCode)
	if errValidCode != nil {
		if errors.Is(errValidCode, rp.ErrProductCodeAlreadyExist) {
			err = errsv.ErrProductConflict{Message: "Product code already exists"}
			return
		}
		err = errsv.ErrProduct{Message: "Error valid product code"}
		return
	}

	if errCreate := p.rp.CreateProduct(&newProduct); errCreate != nil {
		if errors.As(errCreate, &errdb.ErrViolateFK{}) {
			err = errsv.ErrProductConflict{Message: errCreate.Error()}
			return
		}
		err = errsv.ErrProduct{Message: "Error creating product"}
		return
	}

	productDto = product2.ParserProductToDTO(newProduct)
	return
}

func (p ProductService) DeleteProduct(id int) (err error) {
	_, ErrGetEntity := p.GetByID(id)
	if ErrGetEntity != nil {
		err = ErrGetEntity
		return
	}

	errDelete := p.rp.DeleteProduct(id)
	if errDelete != nil {
		if errors.As(errDelete, &errdb.ErrConflict{}) {
			err = errsv.ErrProductConflict{Message: "It is not possible to delete the product because it is being used"}
			return
		}
		err = errsv.ErrProduct{Message: "Error deleted product"}
		return
	}

	return
}

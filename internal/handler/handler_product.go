package handler

import (
	"encoding/json"
	"errors"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	pr "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	product2 "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product"
	"net/http"
	"strconv"
)

func NewProductHandler(sv product.IProductService) *ProductHandle {
	return &ProductHandle{sv: sv}
}

type ProductHandle struct {
	sv product.IProductService
}

func (h *ProductHandle) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, _ := h.sv.GetAll()
		response.JSON(w, http.StatusOK, dto.GenericResponse{Data: products})
	}
}

func (h *ProductHandle) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idPath, errPath := strconv.Atoi(chi.URLParam(r, "id"))
		if errPath != nil {
			response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: "Invalid ID"})
			return
		}

		productSearch, errSearch := h.sv.GetByID(idPath)
		if errSearch != nil {
			validErrorResponse(w, errSearch)
			return
		}

		response.JSON(w, http.StatusOK, dto.GenericResponse{Data: productSearch})
	}
}

func (h *ProductHandle) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idPath, errPath := strconv.Atoi(chi.URLParam(r, "id"))
		if errPath != nil {
			response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: "Invalid ID"})
			return
		}

		errDelete := h.sv.DeleteProduct(idPath)
		if errDelete != nil {
			if errors.As(errDelete, &pr.ErrProductRepository{}) {

				if errors.Is(errDelete, pr.ErrProductNotFound) {
					response.JSON(w, http.StatusNotFound, dto.GenericResponse{Message: errDelete.Error()})
					return
				}

				response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: errDelete.Error()})
				return
			}

			response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: "Internal Server Error"})
			return
		}

		response.JSON(w, http.StatusOK, nil)
	}
}

func (h ProductHandle) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newProduct product2.ProductDTO

		if errDecode := json.NewDecoder(r.Body).Decode(&newProduct); errDecode != nil {
			response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: "Invalid Body"})
			return
		}

		newProduct, errCreate := h.sv.CreateProduct(newProduct)
		if errCreate != nil {

			if errors.As(errCreate, &product.ErrValidProduct{}) {
				response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: errCreate.Error()})
				return
			}

			if errors.As(errCreate, &product.ErrProduct{}) {
				response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: errCreate.Error()})
				return
			}

			response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: "Internal Server Error"})
			return
		}

		response.JSON(w, http.StatusCreated, dto.GenericResponse{Data: newProduct})
	}

}

func (h ProductHandle) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idPath, errPath := strconv.Atoi(chi.URLParam(r, "id"))
		if errPath != nil {
			response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: "Invalid ID"})
			return
		}

		var productRequest product2.UpdateProductRequest
		if errDecode := json.NewDecoder(r.Body).Decode(&productRequest); errDecode != nil {
			response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: "Invalid Body"})
			return
		}

		productUpdate, errPatch := h.sv.UpdateProduct(idPath, productRequest)
		if errPatch != nil {

			if errors.As(errPatch, &product.ErrNotFoundProduct{}) {
				response.JSON(w, http.StatusNotFound, dto.GenericResponse{Message: errPatch.Error()})
				return
			}

			if errors.As(errPatch, &product.ErrValidProduct{}) {
				response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: errPatch.Error()})
				return
			}

			if errors.As(errPatch, &product.ErrProduct{}) {
				response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: errPatch.Error()})
				return
			}

			response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: "Internal Server Error"})
			return
		}

		response.JSON(w, http.StatusOK, dto.GenericResponse{Data: productUpdate})
	}
}

func validErrorResponse(w http.ResponseWriter, err error) {
	if errors.As(err, &product.ErrNotFoundProduct{}) {
		response.JSON(w, http.StatusNotFound, dto.GenericResponse{Message: err.Error()})
		return
	}

	if errors.As(err, &product.ErrValidProduct{}) {
		response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: err.Error()})
		return
	}

	if errors.As(err, &product.ErrProduct{}) {
		response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: err.Error()})
		return
	}

	response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: "Internal Server Error"})
}

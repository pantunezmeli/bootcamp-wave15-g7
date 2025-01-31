package handler

import (
	"encoding/json"
	"errors"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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

func (h *ProductHandle) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idPath, errPath := strconv.Atoi(chi.URLParam(r, "id"))
		if errPath != nil {
			response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: "Invalid ID"})
			return
		}

		errDelete := h.sv.DeleteProduct(idPath)
		if errDelete != nil {
			validErrorResponse(w, errDelete)
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (h ProductHandle) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newProduct product2.ProductDTO

		if errDecode := json.NewDecoder(r.Body).Decode(&newProduct); errDecode != nil {
			response.JSON(w, http.StatusBadRequest, dto.GenericResponse{Message: "Invalid Body"})
			return
		}

		newProduct, errCreate := h.sv.CreateProduct(newProduct)
		if errCreate != nil {
			validErrorResponse(w, errCreate)
			return
		}

		response.JSON(w, http.StatusCreated, dto.GenericResponse{Data: newProduct})
	}

}

func (h ProductHandle) Update() http.HandlerFunc {
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
			validErrorResponse(w, errPath)
			return
		}

		response.JSON(w, http.StatusOK, dto.GenericResponse{Data: productUpdate})

	}
}

func validErrorResponse(w http.ResponseWriter, err error) {

	switch {
	case errors.As(err, &product.ErrNotFoundProduct{}):
		{
			dto.JSONError(w, http.StatusUnprocessableEntity, err.Error())
			break
		}
	case errors.As(err, &product.ErrValidProduct{}):
		{
			dto.JSONError(w, http.StatusUnprocessableEntity, err.Error())
			break
		}
	default:
		{
			//Only for debug
			//fmt.Printf("Error: %v\n", err)
			dto.JSONError(w, http.StatusInternalServerError, dto.ErrInternalServerError.Error())
			break
		}
	}

}

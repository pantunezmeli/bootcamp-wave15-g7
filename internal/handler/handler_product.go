package handler

import (
	"errors"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	pr "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
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
			if errors.Is(errSearch, pr.ErrProductNotFound) {
				response.JSON(w, http.StatusNotFound, dto.GenericResponse{Message: errSearch.Error()})
				return
			}

			response.JSON(w, http.StatusInternalServerError, dto.GenericResponse{Message: "Internal Server Error"})
			return
		}

		response.JSON(w, http.StatusOK, dto.GenericResponse{Data: productSearch})
	}
}

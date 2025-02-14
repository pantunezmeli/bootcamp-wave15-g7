package handler

import (
	"encoding/json"
	"github.com/bootcamp-go/web/response"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product_records"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/productrecords"
	"net/http"
	"strconv"
)

type HandlerProductRecords struct {
	sv product_records.IProductRecordsService
}

func NewHandlerProductRecords(sv product_records.IProductRecordsService) *HandlerProductRecords {
	return &HandlerProductRecords{sv: sv}
}

func (h HandlerProductRecords) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request productrecords.RequestNewRecord
		if errDecode := json.NewDecoder(r.Body).Decode(&request); errDecode != nil {
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidBody.Error())
			return
		}

		newProduct, errCreate := h.sv.CreateProductRecord(request.Data)
		if errCreate != nil {
			validErrorResponse(w, errCreate)
			return
		}

		response.JSON(w, http.StatusCreated, dto.GenericResponse{Data: newProduct})
	}

}

func (h HandlerProductRecords) GetRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idPath := r.URL.Query().Get("id")
		var idParam *int

		if idPath != "" {
			if idInt, err := strconv.Atoi(idPath); err == nil {
				idParam = &idInt
			} else {
				dto.JSONError(w, http.StatusBadRequest, ErrInvalidId.Error())
				return
			}
		}

		productSearch, errSearch := h.sv.GetProductRecord(idParam)
		if errSearch != nil {
			validErrorResponse(w, errSearch)
			return
		}

		response.JSON(w, http.StatusOK, dto.GenericResponse{Data: productSearch})
	}
}

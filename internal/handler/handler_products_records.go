package handler

import (
	"github.com/bootcamp-go/web/response"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product_records"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
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
		//TODO implement me
		panic("implement me")
	}
}

func (h HandlerProductRecords) GetRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var idParam *int
		idPath := r.URL.Query().Get("id")
		if idPath != "" {
			idInt, err := strconv.Atoi(idPath)
			if err != nil {
				dto.JSONError(w, http.StatusBadRequest, ErrInvalidId.Error())
				return
			}
			idParam = &idInt
		} else {
			idParam = nil
		}

		productSearch, errSearch := h.sv.GetProductRecord(idParam)
		if errSearch != nil {
			validErrorResponse(w, errSearch)
			return
		}

		response.JSON(w, http.StatusOK, dto.GenericResponse{Data: productSearch})
	}
}

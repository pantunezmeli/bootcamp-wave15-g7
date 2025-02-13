package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/purchase"
	dtoResponse "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type PurchaseHandler struct {
	service purchase.IServicePurchase
}

func NewPurchaseHandler(sv purchase.IServicePurchase) *PurchaseHandler {
	return &PurchaseHandler{service: sv}
}

func (handler *PurchaseHandler) Get() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		reportPurchases, err := handler.service.GetReport()
		if errors.Is(err, errorbase.ErrEmptyList) {
			dtoResponse.JSONError(writer, http.StatusNotFound, MSG_ErrEmptyList)
			return
		}

		if errors.Is(err, errorbase.ErrDatabaseOperationFailed) {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrOperationDB)
			return
		}

		if err != nil {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrInternalError)
			return
		}
		jsonResponse(writer, http.StatusOK, reportPurchases)

	}
}

func (handler *PurchaseHandler) GetById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		idParam := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrInvalidId)
			return
		}

		reportPurchase, err := handler.service.GetReportById(id)
		if errors.Is(err, errorbase.ErrNotFound) {
			dtoResponse.JSONError(writer, http.StatusNotFound, MSG_ErrNotFound)
			return
		}

		if errors.Is(err, errorbase.ErrDatabaseOperationFailed) {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrOperationDB)
			return
		}

		if err != nil {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrInternalError)
			return
		}
		jsonResponse(writer, http.StatusOK, reportPurchase)

	}
}

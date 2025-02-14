package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/purchase"
	dtoResponse "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/purchase"
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
			dtoResponse.JSONError(writer, http.StatusNoContent, MSG_ErrEmptyList)
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
		if err != nil {
			switch {
			case errors.Is(err, errorbase.ErrNotFound):
				dtoResponse.JSONError(writer, http.StatusNotFound, MSG_ErrNotFound)
			case errors.Is(err, errorbase.ErrDatabaseOperationFailed):
				dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrOperationDB)
			case errors.Is(err, errorbase.ErrInvalidId):
				dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrInvalidId)
			default:
				dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrInternalError)
			}
			return
		}
		jsonResponse(writer, http.StatusOK, reportPurchase)
	}
}

func (controller *PurchaseHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var newPurchase dto.PurchaseOrderResponse
		err := json.NewDecoder(request.Body).Decode(&newPurchase)
		isEmpty := newPurchase == dto.PurchaseOrderResponse{}
		if err != nil || isEmpty {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrJsonFormat)
			return
		}

		purchase, err2 := controller.service.CreatePurchase(newPurchase)
		if err2 != nil {
			switch {
			case errors.Is(err2, errorbase.ErrOrderNumberExist):
				dtoResponse.JSONError(writer, http.StatusConflict, MSG_ErrOrderNumberExist)
			case errors.Is(err2, errorbase.ErrTrackingCodeExist):
				dtoResponse.JSONError(writer, http.StatusConflict, MSG_ErrTrackingCodeExist)
			case errors.Is(err2, errorbase.ErrEmptyParameters),
				errors.Is(err2, errorbase.ErrInvalidRequest):
				dtoResponse.JSONError(writer, http.StatusUnprocessableEntity, MSG_ErrUnprocessable)
			case errors.Is(err2, errorbase.ErrInvalidIdField):
				dtoResponse.JSONError(writer, http.StatusUnprocessableEntity, MSG_ErrInvalidIdField)
			case errors.Is(err2, &errorbase.ErrorFKNotExist{}):
				dtoResponse.JSONError(writer, http.StatusConflict, err2.Error())
			case errors.Is(err2, errorbase.ErrDatabaseOperationFailed):
				dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrOperationDB)
			default:
				dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrInternalError)
			}
			return
		}
		jsonResponse(writer, http.StatusCreated, purchase)
	}
}

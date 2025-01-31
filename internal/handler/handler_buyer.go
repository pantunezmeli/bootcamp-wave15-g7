package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type BuyerHandler struct {
	service buyer.IServiceBuyer
}

func NewBuyerHandler(sv buyer.IServiceBuyer) *BuyerHandler {
	return &BuyerHandler{service: sv}
}

func (handler *BuyerHandler) GetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		buyers, err := handler.service.GetBuyers()
		switch {
		case errors.Is(err, errorbase.ErrEmptyList):
			jsonResponse(writer, http.StatusNotFound, MSG_ErrEmptyList, nil)
			return
		case err != nil:
			jsonResponse(writer, http.StatusInternalServerError, MSG_ErrInternalError, nil)
			return
		default:
			jsonResponse(writer, http.StatusOK, MsgSuccess, buyers)
		}
	}
}

func (handler *BuyerHandler) GetById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		idParam := chi.URLParam(request, "id")
		id, err2 := strconv.Atoi(idParam)

		if err2 != nil {
			jsonResponse(writer, http.StatusInternalServerError, MSG_ErrInternalError, nil)
			return
		}
		buyer, err := handler.service.GetBuyer(id)
		switch {
		case errors.Is(err, errorbase.ErrInvalidId):
			jsonResponse(writer, http.StatusBadRequest, MSG_ErrInvalidId, nil)
			return
		case errors.Is(err, errorbase.ErrNotFound):
			jsonResponse(writer, http.StatusNotFound, MSG_ErrNotFound, nil)
			return
		case err != nil:
			jsonResponse(writer, http.StatusInternalServerError, MSG_ErrInternalError, nil)
			return
		default:
			jsonResponse(writer, http.StatusOK, MsgSuccess, buyer)
		}
	}
}

func (handler *BuyerHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var newBuyer models.Buyer
		err2 := json.NewDecoder(request.Body).Decode(&newBuyer)
		isEmpty := newBuyer == models.Buyer{}

		if err2 != nil || isEmpty {
			jsonResponse(writer, http.StatusBadRequest, MSG_ErrJsonFormat, nil)
			return
		}

		buyer, err := handler.service.CreateBuyer(newBuyer)

		switch {
		case errors.Is(err, errorbase.ErrConflict):
			jsonResponse(writer, http.StatusConflict, MSG_ErrConflict, nil)
			return
		case errors.Is(err, errorbase.ErrStorageOperationFailed):
			jsonResponse(writer, http.StatusInternalServerError, MSG_ErrStorageOperationFailed, nil)
			return
		case err != nil:
			jsonResponse(writer, http.StatusUnprocessableEntity, MSG_ErrIncorrectParameters, nil)
			return
		default:
			jsonResponse(writer, http.StatusCreated, MsgCreated, buyer)
		}
	}

}

func (handler *BuyerHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		idParam := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			jsonResponse(writer, http.StatusBadRequest, MSG_ErrInvalidId, nil)
			return
		}

		var entity dto.BuyerResponse
		if err := json.NewDecoder(request.Body).Decode(&entity); err != nil {
			jsonResponse(writer, http.StatusBadRequest, MSG_ErrJsonFormat, nil)
			return
		}

		buyer, err := handler.service.UpdateBuyer(id, entity)

		switch {
		case errors.Is(err, errorbase.ErrNotFound):
			jsonResponse(writer, http.StatusNotFound, MSG_ErrNotFound, nil)
			return
		case errors.Is(err, errorbase.ErrInvalidId):
			jsonResponse(writer, http.StatusBadRequest, MSG_ErrInvalidId, nil)
			return
		case errors.Is(err, errorbase.ErrStorageOperationFailed):
			jsonResponse(writer, http.StatusInternalServerError, MSG_ErrStorageOperationFailed, nil)
			return
		case err != nil:
			jsonResponse(writer, http.StatusInternalServerError, MSG_ErrInternalError, nil)
			return
		default:
			jsonResponse(writer, http.StatusOK, MsgUpdated, buyer)
		}
	}
}

func (handler *BuyerHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idParam := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			jsonResponse(writer, http.StatusBadRequest, MSG_ErrInvalidId, nil)
			return
		}

		err = handler.service.DeleteBuyer(id)
		switch {
		case errors.Is(err, errorbase.ErrInvalidId):
			jsonResponse(writer, http.StatusBadRequest, MSG_ErrInvalidId, nil)
			return
		case errors.Is(err, errorbase.ErrNotFound):
			jsonResponse(writer, http.StatusNotFound, MSG_ErrNotFound, nil)
			return
		case errors.Is(err, errorbase.ErrStorageOperationFailed):
			jsonResponse(writer, http.StatusInternalServerError, MSG_ErrStorageOperationFailed, nil)
			return
		case err != nil:
			jsonResponse(writer, http.StatusInternalServerError, MSG_ErrInternalError, nil)
			return
		default:
			jsonResponse(writer, http.StatusNoContent, MsgSuccess, nil)
		}
	}
}

func jsonResponse(writer http.ResponseWriter, statusCode int, message string, data any) {
	response.JSON(writer, statusCode, map[string]any{
		"message": message,
		"data":    data,
	})
}

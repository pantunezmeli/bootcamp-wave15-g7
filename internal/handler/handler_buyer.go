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
	dtoResponse "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/buyer"

	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type BuyerHandler struct {
	service buyer.IServiceBuyer
}

func NewBuyerHandler(sv buyer.IServiceBuyer) *BuyerHandler {
	return &BuyerHandler{service: sv}
}

func (handler *BuyerHandler) Get() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		buyers, err := handler.service.GetBuyers()
		if errors.Is(err, errorbase.ErrEmptyList) {
			dtoResponse.JSONError(writer, http.StatusNotFound, MSG_ErrEmptyList)
			return
		}

		if err != nil {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrInternalError)
			return
		}
		jsonResponse(writer, http.StatusOK, buyers)
	}
}

func (handler *BuyerHandler) GetById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		idParam := chi.URLParam(request, "id")
		id, err2 := strconv.Atoi(idParam)

		if err2 != nil {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrRequest)
			return
		}
		buyer, err := handler.service.GetBuyer(id)
		if errors.Is(err, errorbase.ErrInvalidId) {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrInvalidId)
			return
		}
		if errors.Is(err, errorbase.ErrNotFound) {
			dtoResponse.JSONError(writer, http.StatusNotFound, MSG_ErrNotFound)
			return
		}
		if err != nil {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrRequest)
			return

		}
		jsonResponse(writer, http.StatusOK, buyer)
	}
}

func (handler *BuyerHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var newBuyer models.Buyer
		err2 := json.NewDecoder(request.Body).Decode(&newBuyer)
		isEmpty := newBuyer == models.Buyer{}

		if err2 != nil || isEmpty {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrJsonFormat)
			return
		}

		buyer, err := handler.service.CreateBuyer(newBuyer)

		if errors.Is(err, errorbase.ErrConflict) {
			dtoResponse.JSONError(writer, http.StatusConflict, MSG_ErrConflict)
			return
		}

		if errors.Is(err, errorbase.ErrStorageOperationFailed) {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrStorageOperationFailed)
			return
		}

		if err != nil {
			dtoResponse.JSONError(writer, http.StatusConflict, MSG_ErrConflict)
			return
		}

		jsonResponse(writer, http.StatusCreated, buyer)
	}

}

func (handler *BuyerHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		idParam := chi.URLParam(request, "id")
		id, err3 := strconv.Atoi(idParam)

		if err3 != nil {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrInvalidId)
			return
		}

		var entity dto.BuyerResponse
		err2 := json.NewDecoder(request.Body).Decode(&entity)
		if err2 != nil {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrJsonFormat)
			return
		}

		buyer, err := handler.service.UpdateBuyer(id, entity)

		if errors.Is(err, errorbase.ErrNotFound) {
			dtoResponse.JSONError(writer, http.StatusNotFound, MSG_ErrNotFound)
			return
		}

		if errors.Is(err, errorbase.ErrInvalidId) {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrInvalidId)
			return
		}

		if errors.Is(err, errorbase.ErrConflict) {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrConflict)
			return
		}

		if errors.Is(err, errorbase.ErrModelInvalid) {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrModelInvalid)
			return
		}

		if errors.Is(err, errorbase.ErrStorageOperationFailed) {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrStorageOperationFailed)
			return
		}

		if err != nil {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrInternalError)
			return

		}

		jsonResponse(writer, http.StatusOK, buyer)
	}
}

func (handler *BuyerHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idParam := chi.URLParam(request, "id")
		id, err2 := strconv.Atoi(idParam)

		if err2 != nil {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrInvalidId)
			return
		}

		err := handler.service.DeleteBuyer(id)
		if errors.Is(err, errorbase.ErrInvalidId) {
			dtoResponse.JSONError(writer, http.StatusBadRequest, MSG_ErrInvalidId)
			return
		}

		if errors.Is(err, errorbase.ErrNotFound) {
			dtoResponse.JSONError(writer, http.StatusNotFound, MSG_ErrNotFound)
			return
		}

		if errors.Is(err, errorbase.ErrStorageOperationFailed) {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrStorageOperationFailed)
			return
		}

		if err != nil {
			dtoResponse.JSONError(writer, http.StatusInternalServerError, MSG_ErrInternalError)
			return
		}

		jsonResponse(writer, http.StatusNoContent, nil)
	}
}

func jsonResponse(writer http.ResponseWriter, statusCode int, data any) {
	response.JSON(writer, statusCode, map[string]any{
		"data": data,
	})
}

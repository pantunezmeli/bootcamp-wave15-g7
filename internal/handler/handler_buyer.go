package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
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
		if errors.Is(err, errorbase.ErrEmptyList) {
			jsonResponse(writer, http.StatusNotFound, "the buyer list is empty", nil)
			return
		}

		if err != nil {
			jsonResponse(writer, http.StatusBadRequest, "internal server error", nil)
			return
		}
		jsonResponse(writer, http.StatusOK, "success", buyers)
	}
}

func (handler *BuyerHandler) GetBuyerById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		idParam := chi.URLParam(request, "id")
		id, err2 := strconv.Atoi(idParam)

		if err2 != nil {
			jsonResponse(writer, http.StatusInternalServerError, "internal server error", nil)
			return
		}
		buyer, err := handler.service.GetBuyer(id)
		if errors.Is(err, errorbase.ErrInvalidId) {
			jsonResponse(writer, http.StatusBadRequest, "the id parameter is incorrect", nil)
			return
		} else if errors.Is(err, errorbase.ErrNotFound) {
			jsonResponse(writer, http.StatusNotFound, "buyer not found", nil)
			return
		} else if err != nil {
			jsonResponse(writer, http.StatusInternalServerError, "internal server error", nil)
			return
		}
		jsonResponse(writer, http.StatusOK, "success", buyer)
	}
}

func (handler *BuyerHandler) CreateBuyer() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var newBuyer model.Buyer
		err2 := json.NewDecoder(request.Body).Decode(&newBuyer)
		isEmpty := newBuyer == model.Buyer{}

		if err2 != nil || isEmpty {
			jsonResponse(writer, http.StatusBadRequest, "internal server error", nil)
			return
		}

		buyer, err := handler.service.CreateBuyer(newBuyer)

		if errors.Is(err, errorbase.ErrConflict) {
			jsonResponse(writer, http.StatusConflict, "the buyer already exist", nil)
			return
		}

		if err != nil {
			jsonResponse(writer, http.StatusBadRequest, "the fields are empty or incorrect", nil)
			return
		}

		jsonResponse(writer, http.StatusCreated, "buyer created", buyer)
	}

}

func jsonResponse(writer http.ResponseWriter, statusCode int, message string, data any) {
	response.JSON(writer, statusCode, map[string]any{
		"message": message,
		"data":    data,
	})
}

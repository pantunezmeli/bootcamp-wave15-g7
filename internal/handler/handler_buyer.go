package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
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
		if err != nil {
			response.JSON(writer, http.StatusOK, map[string]any{
				"message": http.StatusInternalServerError,
				"data":    nil,
			})
		}

		//list := handler.generateResponseList(buyers)

		response.JSON(writer, http.StatusOK, map[string]any{
			"message": "success",
			"data":    buyers,
		})
	}
}

func (handler *BuyerHandler) GetBuyerById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		idParam := chi.URLParam(request, "id")
		id, err2 := strconv.Atoi(idParam)

		if err2 != nil {
			response.JSON(writer, http.StatusOK, map[string]any{
				"message": http.StatusInternalServerError,
			})
		}

		buyer, err := handler.service.GetBuyer(id)
		if err != nil {
			response.JSON(writer, http.StatusOK, map[string]any{
				"message": http.StatusNotFound,
				"data":    nil,
			})
		}

		//buyerResponse := handler.generateBuyerResponse(buyer)
		response.JSON(writer, http.StatusOK, map[string]any{
			"message": "success",
			"data":    buyer,
		})
	}
}

func (handler *BuyerHandler) CreateBuyer() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var newBuyer model.Buyer
		err2 := json.NewDecoder(request.Body).Decode(&newBuyer)
		isEmpty := newBuyer == model.Buyer{}

		if err2 != nil || isEmpty {
			response.JSON(writer, http.StatusBadRequest, map[string]any{
				"message": http.StatusBadRequest,
			})
			return
		}

		err := handler.service.CreateBuyer(newBuyer)

		if err != nil {
			response.JSON(writer, http.StatusConflict, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		response.JSON(writer, http.StatusCreated, map[string]any{
			"message": http.StatusCreated,
			"data":    "success",
		})
	}

}

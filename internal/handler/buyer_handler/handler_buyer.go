package buyerhandler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	buyerservice "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer_service"
)

type BuyerHandler struct {
	service buyerservice.IServiceBuyer
}

func NewBuyerHandler(sv buyerservice.IServiceBuyer) *BuyerHandler {
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

		list := handler.generateResponseList(buyers)

		response.JSON(writer, http.StatusOK, map[string]any{
			"message": "success",
			"data":    list,
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

		buyerResponse := handler.generateBuyerResponse(buyer)
		response.JSON(writer, http.StatusOK, map[string]any{
			"message": "success",
			"data":    buyerResponse,
		})
	}
}

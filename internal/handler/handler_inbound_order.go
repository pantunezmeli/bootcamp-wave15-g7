package handler

import (
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/inboundorder"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/inboundorder"
)

var ErrOrderNumberExists error = errors.New("order number already exists")
var ErrProductBatchNotFound = errors.New("product batch not found")

func NewInboundOrderHandler(service sv.InboundOrderService) *InboundOrderHandler {
	return &InboundOrderHandler{sv: service}
}

type InboundOrderHandler struct {
	sv sv.InboundOrderService
}

func (h *InboundOrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var inboundOrderData inboundorder.InboundOrderDoc
		if err := request.JSON(r, &inboundOrderData); err != nil {
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidBody.Error())
			return
		}

		// process
		newInboundOrder, err := h.sv.New(inboundOrderData)
		if err != nil {
			h.handleError(w, err)
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": newInboundOrder,
		})
	}
}

func (h *InboundOrderHandler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, sv.ErrEmptyField):
		dto.JSONError(w, http.StatusUnprocessableEntity, ErrEmptyField.Error())
	case errors.Is(err, sv.ErrOrderNumberAlreadyExists):
		dto.JSONError(w, http.StatusConflict, ErrOrderNumberExists.Error())
	case errors.Is(err, sv.ErrEmployeeNotFound):
		dto.JSONError(w, http.StatusNotFound, ErrEmployeeNotFound.Error())
	case errors.Is(err, sv.ErrProductBatchNotFound):
		dto.JSONError(w, http.StatusNotFound, ErrProductBatchNotFound.Error())
	case errors.Is(err, sv.ErrWarehouseNotFound):
		dto.JSONError(w, http.StatusNotFound, ErrWarehouseNotFound.Error())
	default:
		dto.JSONError(w, http.StatusInternalServerError, ErrInternalServerError.Error())
	}
}

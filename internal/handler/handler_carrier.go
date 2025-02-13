package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	service "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/carrier"
	e "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/carrier"
)

type CarrierHandler struct {
	sv service.ICarrierService
}

func NewCarrierHandler(sv service.ICarrierService) *CarrierHandler {
	return &CarrierHandler{sv: sv}
}

// ! 1)
func (h *CarrierHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CarrierDoc
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			e.JSONError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		if req.Cid == "" || req.Address == "" || req.Telephone == "" ||
			req.LocalityId <= 0 || req.CompanyName == "" {
			e.JSONError(w, http.StatusUnprocessableEntity, "missing or invalid required fields")
			return
		}

		carrier, err := h.sv.AddCarrier(req)
		if err != nil {
			var errFK service.ErrForeignKey
			var errDuplicate service.ErrDuplicate
			var errValidation service.ErrInvalidParameter
			var errDB service.ErrDatabase
			var errDTO service.ErrConvertion

			switch {
			case errors.As(err, &errValidation):
				e.JSONError(w, http.StatusBadRequest, "some values are not valid")
				return
			case errors.As(err, &errFK):
				e.JSONError(w, http.StatusNotFound, "locality does not exists")
				return
			case errors.As(err, &errDuplicate):
				e.JSONError(w, http.StatusConflict, fmt.Sprintf("carrier with %s cid already exists", req.Cid))
				return
			case errors.As(err, &errDB):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			case errors.As(err, &errDTO):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			default:
				e.JSONError(w, http.StatusInternalServerError, "unexpected error, try again later")
				return
			}
		}

		response.JSON(w, http.StatusCreated, carrier)
	}
}

// ! 2)
func (h *CarrierHandler) GetCarriesAmount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			e.JSONError(w, http.StatusBadRequest, "missing locality ID")
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			e.JSONError(w, http.StatusBadRequest, "invalid id parameter")
			return
		}

		res, err := h.sv.GetCarriesAmount(id)
		if err != nil {
			var errDB service.ErrDatabase
			var errDTO service.ErrConvertion
			switch {
			case errors.As(err, &errDB):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			case errors.As(err, &errDTO):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			case errors.Is(err, service.ErrLocalityNotExists):
				e.JSONError(w, http.StatusNotFound, "locality not found")
				return
			default:
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			}
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": res,
		})
	}
}

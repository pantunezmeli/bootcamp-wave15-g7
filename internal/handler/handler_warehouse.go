package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	customErrors "github.com/pantunezmeli/bootcamp-wave15-g7/internal/errors"
	service "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse"
	e "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/warehouse"
)

type WareHouseHandler struct {
	sv service.IWareHouseService
}

func NewWareHouseHandler(sv service.IWareHouseService) *WareHouseHandler {
	return &WareHouseHandler{sv: sv}
}

// ! 1)
func (h *WareHouseHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wh, err := h.sv.FindAll()
		if err != nil {
			var errMap customErrors.ErrConvertion
			var errDB customErrors.ErrDatabase

			switch {
			case errors.As(err, &errMap):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			case errors.As(err, &errDB):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			default:
				e.JSONError(w, http.StatusInternalServerError, "unexpected error, try again later")
				return
			}
		}

		var whList []dto.WareHouseDoc
		for _, warehouse := range wh {
			whList = append(whList, warehouse)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": whList,
		})
	}
}

// ! 2)
func (h *WareHouseHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			e.JSONError(w, http.StatusBadRequest, "invalid id parameter")
			return
		}

		wh, err := h.sv.GetWareHouseById(id)

		if err != nil {
			var errNotFound customErrors.ErrNotFound
			var errMap customErrors.ErrConvertion
			var errDB customErrors.ErrDatabase

			switch {
			case errors.As(err, &errNotFound):
				e.JSONError(w, http.StatusNotFound, fmt.Sprintf("warehouse with id %d not found", id))
				return
			case errors.As(err, &errMap):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			case errors.As(err, &errDB):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			default:
				e.JSONError(w, http.StatusInternalServerError, "unexpected error, try again later")
				return
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": wh,
		})
	}
}

// ! 3)
func (h *WareHouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req dto.WareHouseDoc

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			e.JSONError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		if req.WareHouseCode == "" || req.Address == "" || req.Telephone == "" ||
			req.LocalityId <= 0 {
			e.JSONError(w, http.StatusUnprocessableEntity, "missing or invalid required fields")
			return
		}

		wh, err := h.sv.AddWareHouse(req)
		if err != nil {
			var errFK customErrors.ErrForeignKey
			var errDuplicate customErrors.ErrDuplicate
			var errValidation customErrors.ErrInvalidParameter
			var errDB customErrors.ErrDatabase
			var errDTO customErrors.ErrConvertion

			switch {
			case errors.As(err, &errValidation):
				e.JSONError(w, http.StatusBadRequest, "some values are not valid")
				return
			case errors.As(err, &errFK):
				e.JSONError(w, http.StatusNotFound, "locality does not exists")
				return
			case errors.As(err, &errDuplicate):
				e.JSONError(w, http.StatusConflict, fmt.Sprintf("warehouse with warehouse code %s already exists", req.WareHouseCode))
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

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": wh,
		})

	}
}

// ! 4)
func (h *WareHouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")

		var req dto.WareHouseDoc

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			e.JSONError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			e.JSONError(w, http.StatusBadRequest, "invalid id parameter")
			return
		}

		wh, err := h.sv.EditWareHouse(id, req)

		if err != nil {
			var errFK customErrors.ErrForeignKey
			var errDuplicate customErrors.ErrDuplicate
			var errNotFound customErrors.ErrNotFound
			var errDB customErrors.ErrDatabase
			var errDTO customErrors.ErrConvertion

			switch {
			case errors.As(err, &errNotFound):
				e.JSONError(w, http.StatusNotFound, "warehouse not found")
				return
			case errors.As(err, &errFK):
				e.JSONError(w, http.StatusConflict, "locality does not exists")
				return
			case errors.As(err, &errDuplicate):
				e.JSONError(w, http.StatusConflict, fmt.Sprintf("warehouse with warehouse code %s already exists", req.WareHouseCode))
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

		response.JSON(w, http.StatusOK, map[string]any{
			"data": wh,
		})

	}
}

// ! 5)
func (h *WareHouseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			e.JSONError(w, http.StatusBadRequest, "invalid id paramete")
			return
		}

		err = h.sv.DeleteWarehouse(id)

		if err != nil {
			var errNotFound customErrors.ErrNotFound
			var errDB customErrors.ErrDatabase

			switch {
			case errors.As(err, &errNotFound):
				e.JSONError(w, http.StatusNotFound, "warehouse not found")
				return
			case errors.As(err, &errDB):
				e.JSONError(w, http.StatusInternalServerError, "something went wrong, try again later")
				return
			default:
				e.JSONError(w, http.StatusInternalServerError, "unexpected error, try again later")
				return
			}
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

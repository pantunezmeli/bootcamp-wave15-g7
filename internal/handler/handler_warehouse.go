package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	service "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse"
	e "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/warehouse"
)

type WareHouseHandler struct {
	sv service.IWareHouseService
}

// Constructor
func NewWareHouseHandler(sv service.IWareHouseService) *WareHouseHandler {
	return &WareHouseHandler{sv: sv}
}

// ! 1)
func (h *WareHouseHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wh, err := h.sv.FindAll()
		if err != nil {
			e.JSONError(w, http.StatusInternalServerError, "please try again later")
			return
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
			if errors.Is(err, service.ErrWareHouseNotFound) {
				e.JSONError(w, http.StatusNotFound, "warehouse not found")
				return
			}
			e.JSONError(w, http.StatusInternalServerError, "please try again later")
			return
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
			req.MinimunCapacity <= 0 || req.MinimunTemperature < -100 {
			e.JSONError(w, http.StatusUnprocessableEntity, "missing or invalid required fields")
			return
		}

		wh, err := h.sv.AddWareHouse(req)

		if err != nil {
			if errors.Is(err, service.ErrWareHouseCodeAlreadyExists) {
				e.JSONError(w, http.StatusConflict, "warehouse with warehouse_code already exists")
				return
			}

			var invalidFieldErr *service.ErrInvalidParameter
			if errors.As(err, &invalidFieldErr) {
				e.JSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			e.JSONError(w, http.StatusInternalServerError, "please try again later")
			return
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
			if errors.Is(err, service.ErrWareHouseNotFound) {
				e.JSONError(w, http.StatusNotFound, "warehouse not found")
				return
			}

			if errors.Is(err, service.ErrWareHouseCodeAlreadyExists) {
				e.JSONError(w, http.StatusConflict, "warehouse with taht code already exists")
				return
			}

			var invalidFieldErr *service.ErrInvalidParameter
			if errors.As(err, &invalidFieldErr) {
				e.JSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			e.JSONError(w, http.StatusInternalServerError, "please try again later")
			return
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
			if errors.Is(err, service.ErrWareHouseNotFound) {
				e.JSONError(w, http.StatusNotFound, "warehouse not found")
				return
			}
			e.JSONError(w, http.StatusInternalServerError, "please try again later")
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

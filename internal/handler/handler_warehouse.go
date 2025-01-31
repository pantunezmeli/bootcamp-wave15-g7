package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	service "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse_service"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
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

		// Call Service
		wh, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// Convert map to list of maps
		var whList []dto.WareHouseDoc
		for _, warehouse := range wh {
			whList = append(whList, warehouse)
		}

		// Right Response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    whList,
		})
	}
}

// ! 2)
func (h *WareHouseHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Params
		idStr := chi.URLParam(r, "id")

		// Param validation
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id parameter",
			})
			return
		}

		// Call service
		wh, err := h.sv.GetWareHouseById(id)

		// Errors types
		if err != nil {
			if errors.Is(err, service.ErrWareHouseNotFound) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": err.Error(),
				})
				return
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// Right response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    wh,
		})
	}
}

// ! 3)
func (h *WareHouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req dto.WareHouseDoc

		// Decode request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request payload",
			})
			return
		}

		// Validation of required camps
		if req.WareHouseCode == "" || req.Address == "" || req.Telephone == "" ||
			req.MinimunCapacity <= 0 || req.MinimunTemperature < -100 {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "missing or invalid required fields",
				"errors": map[string]string{
					"warehouse_code":      "must not be empty",
					"address":             "must not be empty",
					"telephone":           "must not be empty",
					"minimun_capacity":    "must be greater than 0",
					"minimun_temperature": "must be at least -100",
				},
			})
			return
		}

		// Call Service
		wh, err := h.sv.AddWareHouse(req)

		// Errors types
		if err != nil {
			if errors.Is(err, service.ErrWareHouseCodeAlreadyExists) {
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "warehouse with warehouse_code already exists",
				})
				return
			}

			var invalidFieldErr error
			if errors.As(err, &invalidFieldErr) {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": "some values are not valid",
				})
				return
			}

			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "something went wrong",
			})
			return
		}

		// Right Response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    wh,
		})

	}
}

// ! 4)
func (h *WareHouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Params
		idStr := chi.URLParam(r, "id")

		var req dto.WareHouseDoc

		// Decode request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request payload",
			})
			return
		}

		// Param validation
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id parameter",
			})
			return
		}

		// Call service
		wh, err := h.sv.EditWareHouse(id, req)

		// Errors types
		if err != nil {
			if errors.Is(err, service.ErrWareHouseNotFound) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "warehouse not found",
				})
				return
			}

			if errors.Is(err, service.ErrWareHouseCodeAlreadyExists) {
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "warehouse with taht code already exists",
				})
				return
			}

			var invalidFieldErr error
			if errors.As(err, &invalidFieldErr) {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": "some values are not valid",
				})
				return
			}

			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "something went wrong",
			})
			return
		}

		// Right response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    wh,
		})

	}
}

// ! 5)
func (h *WareHouseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Params
		idStr := chi.URLParam(r, "id")

		// Param validation
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id parameter",
			})
			return
		}

		// Call Service
		err = h.sv.DeleteWarehouse(id)

		// Errors types
		if err != nil {
			if errors.Is(err, service.ErrWareHouseNotFound) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "warehouse not found",
				})
				return
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "something went wrong",
			})
			return
		}

		// Right response
		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "warehouse eliminated successfully",
		})
	}
}

package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	service "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse_service"
)

type WareHouseHandler struct {
	sv service.IWareHouseService
}

// Constructor
func NewWareHouseHandler(sv service.IWareHouseService) *WareHouseHandler {
	return &WareHouseHandler{sv: sv}
}

func (h *WareHouseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wh, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		print("Este es el wh: \n", wh)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    wh,
		})
	}
}

func (h *WareHouseHandler) GetWareHouseById() http.HandlerFunc {
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

		// Any error
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

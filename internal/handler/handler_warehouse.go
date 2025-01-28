package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
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

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    wh,
		})
	}
}

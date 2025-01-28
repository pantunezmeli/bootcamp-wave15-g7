package product

import (
	"github.com/bootcamp-go/web/response"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product"
	"net/http"
)

func NewProductHandler(sv product.IProductService) *ProductHandle {
	return &ProductHandle{sv: sv}
}

type ProductHandle struct {
	sv product.IProductService
}

func (h *ProductHandle) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

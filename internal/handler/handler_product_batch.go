package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	productbatch "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product_batch"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product_batch"
)

// ProductBatchHandler is a struct that represents a ProductBatch handler
type ProductBatchHandler struct {
	service productbatch.IProductBatchService
}

// NewProductBatchHandler is a function that returns a new instance of ProductBatchHandler
func NewProductBatchHandler(sv productbatch.IProductBatchService) *ProductBatchHandler {
	return &ProductBatchHandler{service: sv}
}

// Create is a method that creates a new ProductBatch
func (h *ProductBatchHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var productBatchDto dto.ProductBatchResponse
		err := json.NewDecoder(r.Body).Decode(&productBatchDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		productBatchResponse, err := h.service.CreateProductBatch(productBatchDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": productBatchResponse,
		})
	}
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	productbatch "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product_batch"
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
		var productBatch models.ProductBatch
		err := json.NewDecoder(r.Body).Decode(&productBatch)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		productBatchResponse, err := h.service.CreateProductBatch(productBatch)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(productBatchResponse)
		json.NewEncoder(w).Encode(productBatch)
	}
}

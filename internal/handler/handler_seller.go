package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/seller"
)

type SellerDefault struct {
	sv seller.SellerService
}

func NewSellerDefault(sv seller.SellerService) *SellerDefault{
	return &SellerDefault{sv}
}


func (h *SellerDefault) GetAll() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		res, err := h.sv.GetAll()
		if err != nil {
			switch err {
			default:
				response.Error(w, http.StatusInternalServerError, "please try again later")
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": res,
		})
	}
}
package handler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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
			switch {
			default:
				response.Error(w, http.StatusInternalServerError, "please try again later")
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": res,
		})
	}
}

func (h *SellerDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}
		idParsed, err := strconv.Atoi(id)
		if err != nil{
			response.Error(w, http.StatusBadRequest, "id should be a number")
		}
		res, err := h.sv.GetById(idParsed)
		if err != nil {
			switch {
				
			}
		}
	}
}
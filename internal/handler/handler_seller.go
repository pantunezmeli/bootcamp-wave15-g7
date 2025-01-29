package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	repo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/seller"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

// cambiar errores
// Preguntar si response es un array 

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
			return
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
			return
		}
		res, err := h.sv.GetById(idParsed)
		if err != nil {
			switch {
			case errors.Is(err, repo.ErrSellerNotFound):
				response.Error(w, http.StatusNotFound, "seller not found")
			default:
				response.Error(w, http.StatusInternalServerError, "please try again later")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": res,
		})
	}
}


func (h *SellerDefault) Create() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody dto.SellerDoc
		if err := request.JSON(r, &reqBody); err != nil{
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		res, err := h.sv.Save(reqBody)
		if err != nil {
			//handlear los errores que faltan
			switch{
			case errors.Is(err, repo.ErrCidAlreadyExists):
				response.Error(w, http.StatusConflict, "cid already exists")
			default:
				response.Error(w, http.StatusInternalServerError, "please try again later")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": res,
		})

	}
}

func (h *SellerDefault) Delete() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//CODIGO REPETIDO
		id := chi.URLParam(r, "id")
		if id == "" {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}
		idParsed, err := strconv.Atoi(id)
		if err != nil{
			response.Error(w, http.StatusBadRequest, "id should be a number")
			return
		}

		err = h.sv.Delete(idParsed)
		if err != nil {
			switch {
			case errors.Is(err, repo.ErrSellerNotFound):
				response.Error(w, http.StatusNotFound, "seller not found")
			default:
				response.Error(w, http.StatusInternalServerError, "please try again later")
			}
			return
		}
		response.Text(w, http.StatusNoContent, "seller deleted")

	}
}
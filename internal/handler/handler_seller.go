package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	repo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/seller"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	seller_dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/seller"
)


var (
	ErrInternalServerError = errors.New("internal server error, please try again later")
	ErrInvalidId = errors.New("invalid id, id should be a number")
	ErrSellerNotFound = errors.New("seller not found")
	ErrInvalidBody = errors.New("invalid body")
	ErrCidExists = errors.New("cid already exists and should be unique")
	ErrLocalityNotExist = errors.New("locality_id does not exist")
)


type SellerDefault struct {
	sv seller.SellerService
}

func NewSellerDefault(sv seller.SellerService) *SellerDefault{
	return &SellerDefault{sv}
}


func (h *SellerDefault) Get() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		res, err := h.sv.GetAll()
		if err != nil {
			dto.JSONError(w, http.StatusInternalServerError, ErrInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": res,
		})
	}
}

func (h *SellerDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParsed, isInvalid := validateId(r, w)
		if isInvalid {
			return
		}

		res, err := h.sv.GetById(idParsed)
		if err != nil {
			switch {
			case errors.Is(err, repo.ErrSellerNotFound):
				dto.JSONError(w, http.StatusNotFound, ErrSellerNotFound.Error())
			default:
				dto.JSONError(w, http.StatusInternalServerError, ErrInternalServerError.Error())
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
		var reqBody seller_dto.SellerRequest
		if err := request.JSON(r, &reqBody); err != nil{
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidBody.Error())
			return
		}

		res, err := h.sv.Save(reqBody)
		if err != nil {
			var missingParamErr *seller.ErrMissingParameters
			var invalidParamErr *seller.ErrInvalidParameter
			switch{
			case errors.Is(err, repo.ErrCidAlreadyExists):
				dto.JSONError(w, http.StatusConflict, ErrCidExists.Error())
			case errors.Is(err, repo.ErrLocalityNotFound):
				dto.JSONError(w, http.StatusConflict, ErrLocalityNotExist.Error())
			case errors.As(err, &missingParamErr):
				dto.JSONError(w, http.StatusUnprocessableEntity, fmt.Sprintf("missing parameter: %s", missingParamErr.Error()))
			case errors.As(err, &invalidParamErr):
				dto.JSONError(w, http.StatusUnprocessableEntity, fmt.Sprintf("invalid parameter: %s", invalidParamErr.Error()))
			default:
				dto.JSONError(w, http.StatusInternalServerError, ErrInternalServerError.Error())
			}
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": res,
		})

	}
}

func (h *SellerDefault) Delete() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idParsed, isInvalid := validateId(r, w)
		if isInvalid {
			return
		}

		err := h.sv.Delete(idParsed)
		if err != nil {
			switch {
			case errors.Is(err, repo.ErrSellerNotFound):
				dto.JSONError(w, http.StatusNotFound, ErrSellerNotFound.Error())
			default:
				dto.JSONError(w, http.StatusInternalServerError, ErrInternalServerError.Error())
			}
			return
		}

		response.JSON(w, http.StatusNoContent, nil)

	}
}


func (h *SellerDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParsed, isInvalid := validateId(r, w)
		if isInvalid {
			return
		}

		var reqBody seller_dto.SellerRequest
		if err := request.JSON(r, &reqBody); err != nil{
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidBody.Error())
			return
		}

		res, err := h.sv.Update(reqBody, idParsed)
		if err != nil {
			var invalidParamErr *seller.ErrInvalidParameter
			switch{
			case errors.Is(err, repo.ErrCidAlreadyExists):
				dto.JSONError(w, http.StatusConflict, ErrCidExists.Error())
			case errors.Is(err, repo.ErrSellerNotFound):
				dto.JSONError(w, http.StatusNotFound, ErrSellerNotFound.Error())
			case errors.Is(err, repo.ErrLocalityNotFound):
				dto.JSONError(w, http.StatusConflict, ErrLocalityNotExist.Error())
			case errors.As(err, &invalidParamErr):
				dto.JSONError(w, http.StatusUnprocessableEntity, fmt.Sprintf("invalid parameter: %s", invalidParamErr.Error()))
			default:
				dto.JSONError(w, http.StatusInternalServerError, ErrInternalServerError.Error())
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": res,
		})



	}
}

func validateId(r *http.Request, w http.ResponseWriter) (int, bool) {
	id := chi.URLParam(r, "id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		dto.JSONError(w, http.StatusBadRequest, ErrInvalidId.Error())
		return 0, true
	}
	return idParsed, false
}

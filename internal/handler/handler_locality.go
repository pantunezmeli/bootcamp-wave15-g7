package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	repository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/locality"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/locality"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	locality_dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/locality"
)

type LocalityDefault struct {
	sv locality.LocalityService
}

func NewLocalityDefault(sv locality.LocalityService) *LocalityDefault {
	return &LocalityDefault{sv}
}

func (h *LocalityDefault) Create() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody locality_dto.LocalityRequest
		if err := request.JSON(r, &reqBody); err != nil {
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidBody.Error())
			return
		}

		res, err := h.sv.Save(reqBody)
		if err != nil {
			handleLocalityError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any {
			"data": res,
		})
	}
}

func (h *LocalityDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParsed, isInvalid := validateId(r, w)
		if isInvalid {
			return
		}

		res, err := h.sv.GetById(idParsed)
		if err != nil {
			handleLocalityError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": res,
		})


	}
}

func (h *LocalityDefault) GetReportSellers() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var id *int
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			parsedID, err := strconv.Atoi(idStr)
			if err != nil{
				dto.JSONError(w, http.StatusBadRequest, ErrInvalidId.Error())
				return
			}
			id = &parsedID
		}

		res, err := h.sv.GetReportSellers(id)

		if err != nil {
			handleLocalityError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": res,
		})

		
	}
}

func handleLocalityError(w http.ResponseWriter, err error) {
	errorMap := map[error]int{
		repository.ErrProvinceNotFound: http.StatusConflict,
		repository.ErrLocalityNotFound: http.StatusNotFound,
	}

	for key, status := range errorMap {
		if errors.Is(err, key) {
			dto.JSONError(w, status, key.Error())
			return
		}
	}

	var missingParamErr *locality.ErrMissingParameters
	var invalidParamErr *locality.ErrInvalidParameter

	switch {
	case errors.As(err, &missingParamErr):
		dto.JSONError(w, http.StatusUnprocessableEntity, fmt.Sprintf("missing parameter: %s", missingParamErr.Error()))
	case errors.As(err, &invalidParamErr):
		dto.JSONError(w, http.StatusUnprocessableEntity, fmt.Sprintf("invalid parameter: %s", invalidParamErr.Error()))
	default:
		dto.JSONError(w, http.StatusInternalServerError, ErrInternalServerError.Error())
	}
}

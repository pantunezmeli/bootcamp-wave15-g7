package handler

import (
	"encoding/json"
	"net/http"

	"strconv"

	"errors"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/section"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	dtoSection "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/section"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

// SectionDefault is a struct with methods that represent handlers for sections
type SectionDefault struct {
	// service is the service that will be used by the handler
	service section.ISectionService
}

// NewSectionDefault is a function that returns a new instance of SectionDefault
func NewSectionDefault(sv section.ISectionService) *SectionDefault {
	return &SectionDefault{service: sv}
}

// ListSections is a method that returns a handler for the route GET /sections
func (h *SectionDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all Sections
		sections, err := h.service.ListSections()
		if err != nil {
			dto.JSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": sections,
		})
	}
}

func (h *SectionDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtein the ID from the URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			dto.JSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Get the Section by ID
		section, err := h.service.GetSection(id)
		if err != nil {
			if errors.Is(err, errorbase.ErrNotFound) {
				dto.JSONError(w, http.StatusNotFound, err.Error())
				return
			}
			if errors.Is(err, errorbase.ErrInvalidId) {
				dto.JSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			dto.JSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": section,
		})
	}
}

func (h *SectionDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode the request body
		var newSection models.Section
		if err := json.NewDecoder(r.Body).Decode(&newSection); err != nil {
			dto.JSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Call Service to Create Section
		section, err := h.service.CreateSection(newSection)

		// Handle Errors
		if err != nil {
			if errors.Is(err, errorbase.ErrConflict) {
				dto.JSONError(w, http.StatusConflict, err.Error())
				return
			}
			if errors.Is(err, errorbase.ErrEmptyParameters) {
				dto.JSONError(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
			if errors.Is(err, errorbase.ErrInvalidRequest) {
				dto.JSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			dto.JSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Response
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": section,
		})
	}
}

func (h *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the ID from the URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			dto.JSONError(w, http.StatusBadRequest, errorbase.ErrInvalidId.Error())
			return
		}

		// Decode the request body
		var sectionDoc dtoSection.SectionResponse
		if err := json.NewDecoder(r.Body).Decode(&sectionDoc); err != nil {
			dto.JSONError(w, http.StatusBadRequest, errorbase.ErrInvalidRequest.Error())
			return
		}

		section, err := h.service.PatchSection(id, sectionDoc)

		// Handle Errors
		if err != nil {
			if errors.Is(err, errorbase.ErrNotFound) {
				dto.JSONError(w, http.StatusNotFound, err.Error())
				return
			}
			// if errors.Is(err, errorbase.ErrConflict) {
			// 	dto.JSONError(w, http.StatusConflict, err.Error())
			// 	return
			// }

			dto.JSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": section,
		})
	}
}

func (h *SectionDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener la ID de la URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			dto.JSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Get the Section by ID
		err = h.service.DeleteSection(id)
		if err != nil {
			if errors.Is(err, errorbase.ErrNotFound) {
				dto.JSONError(w, http.StatusNotFound, err.Error())
				return
			}
			if errors.Is(err, errorbase.ErrInvalidId) {
				dto.JSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			dto.JSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Response
		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "section deleted succesfully",
		})
	}
}

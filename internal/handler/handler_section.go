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
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/errorbase"
)

// SectionDefault is a struct with methods that represent handlers for sections
type SectionDefault struct {
	// sv is the service that will be used by the handler
	service section.ISectionService
}

// NewSectionDefault is a function that returns a new instance of SectionDefault
func NewSectionDefault(sv section.ISectionService) *SectionDefault {
	return &SectionDefault{service: sv}
}

// ListSections is a method that returns a handler for the route GET /sections
func (h *SectionDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// process
		// - get all Sections
		sections, err := h.service.ListSections()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": sections,
		})
	}
}

func (h *SectionDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener la ID de la URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
			return
		}

		// process
		// - get the Section by ID
		section, err := h.service.GetSection(id)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

		// response
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
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		// Call Service to Create Section
		section, err := h.service.CreateSection(newSection)

		// Handle Errors
		if errors.Is(err, errorbase.ErrConflict) {
			response.JSON(w, http.StatusConflict, nil)
			return
		}
		if errors.Is(err, errorbase.ErrStorageOperationFailed) {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, nil)
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
		// Obtener la ID de la URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": errorbase.ErrInvalidId.Error()})
			return
		}

		// Decodificar el cuerpo de la solicitud
		var sectionDoc dto.SectionResponse
		if err := json.NewDecoder(r.Body).Decode(&sectionDoc); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": errorbase.ErrInvalidRequest.Error()})
			return
		}

		section, err := h.service.PatchSection(id, sectionDoc)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

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
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
			return
		}

		// process
		// - get the Section by ID
		err = h.service.DeleteSection(id)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "section deleted succesfully",
		})
	}
}

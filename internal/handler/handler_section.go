package handler

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/section"
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
		v, err := h.service.ListSections()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

		// response
		data := make(map[int]models.SectionDoc)
		for key, value := range v {
			data[key] = ConvertSectionToSectionDoc(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
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
		data := ConvertSectionToSectionDoc(section)
		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *SectionDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v models.SectionDoc
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		section := ConvertSectionDocToSection(v)

		if err := h.service.CreateSection(section); err != nil {
			if err.Error() == "section already exists" {
				response.JSON(w, http.StatusConflict, nil)
			} else {
				response.JSON(w, http.StatusInternalServerError, nil)
			}
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "section created successfully",
		})
	}
}

func (h *SectionDefault) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener la ID de la URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": errorbase.ErrInvalidId.Error()})
			return
		}
		// Decodificar el cuerpo de la solicitud
		var sectionDoc models.SectionDoc
		if err := json.NewDecoder(r.Body).Decode(&sectionDoc); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": errorbase.ErrInvalidRequest.Error()})
			return
		}

		section := ConvertSectionDocToSection(sectionDoc)

		section, err = h.service.PatchSection(id, section)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

		// response
		data := ConvertSectionToSectionDoc(section)
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
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

// ConvertSectionDocToSection converts a SectionDoc to a Section
func ConvertSectionDocToSection(vd models.SectionDoc) models.Section {
	return models.Section{
		Id:                  vd.Id,
		Section_Number:      vd.Section_Number,
		Current_Temperature: vd.Current_Temperature,
		Minimum_Temperature: vd.Minimum_Temperature,
		Current_Capacity:    vd.Current_Capacity,
		Minimum_Capacity:    vd.Minimum_Capacity,
		Maximum_Capacity:    vd.Maximum_Capacity,
		Warehouse_Id:        vd.Warehouse_Id,
		Product_Type_Id:     vd.Product_Type_Id,
	}
}

// ConvertSectionToSectionDoc convierte un Section a un SectionDoc
func ConvertSectionToSectionDoc(v models.Section) models.SectionDoc {
	return models.SectionDoc{
		Id:                  v.Id,
		Section_Number:      v.Section_Number,
		Current_Temperature: v.Current_Temperature,
		Minimum_Temperature: v.Minimum_Temperature,
		Current_Capacity:    v.Current_Capacity,
		Minimum_Capacity:    v.Minimum_Capacity,
		Maximum_Capacity:    v.Maximum_Capacity,
		Warehouse_Id:        v.Warehouse_Id,
		Product_Type_Id:     v.Product_Type_Id,
	}
}

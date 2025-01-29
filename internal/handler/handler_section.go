package handler

import (
	"net/http"

	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/section"
)

// NewSectionDefault is a function that returns a new instance of SectionDefault
func NewSectionDefault(sv section.SectionServiceV2) *SectionDefault {
	return &SectionDefault{sv: sv}
}

// SectionDefault is a struct with methods that represent handlers for sections
type SectionDefault struct {
	// sv is the service that will be used by the handler
	sv section.SectionServiceV2
}

// ListSections is a method that returns a handler for the route GET /sections
func (h *SectionDefault) ListSections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all Sections
		v, err := h.sv.ListSections()
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
			"message": "success",
			"data":    data,
		})
	}
}

func (h *SectionDefault) GetSection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener la ID de la URL
		idStr := chi.URLParam(r, "id")

		// Obtener la ID de la URL
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
			return
		}

		// process
		// - get the Section by ID
		section, err := h.sv.GetSection(id)
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

func (h *SectionDefault) CreateSection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - convert SectionDoc to Section
		// - create Section
		// - response
	}
}

func (h *SectionDefault) PatchSection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - convert SectionDoc to Section
		// - update Section
		// - response
	}
}

func (h *SectionDefault) DeleteSection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - delete Section
		// - response
	}
}

// ConvertSectionDocToSection converts a SectionDoc to a Section
func ConvertSectionDocToSection(vd models.SectionDoc) models.Section {
	return models.Section{
		Id:                  vd.Id,
		Section_number:      vd.Section_number,
		Current_temperature: vd.Current_temperature,
		Minimum_temperature: vd.Minimum_temperature,
		Current_capacity:    vd.Current_capacity,
		Minimum_capacity:    vd.Minimum_capacity,
		Maximim_capacity:    vd.Maximim_capacity,
		Warehouse_id:        vd.Warehouse_id,
		Product_type_id:     vd.Product_type_id,
	}
}

// ConvertSectionToSectionDoc convierte un Section a un SectionDoc
func ConvertSectionToSectionDoc(v models.Section) models.SectionDoc {
	return models.SectionDoc{
		Id:                  v.Id,
		Section_number:      v.Section_number,
		Current_temperature: v.Current_temperature,
		Minimum_temperature: v.Minimum_temperature,
		Current_capacity:    v.Current_capacity,
		Minimum_capacity:    v.Minimum_capacity,
		Maximim_capacity:    v.Maximim_capacity,
		Warehouse_id:        v.Warehouse_id,
		Product_type_id:     v.Product_type_id,
	}
}

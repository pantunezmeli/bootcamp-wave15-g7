package handler

import (
	"encoding/json"
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
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
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
		var v models.SectionDoc
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		section := ConvertSectionDocToSection(v)

		if err := h.sv.CreateSection(section); err != nil {
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

// func (h *SectionDefault) PatchSection() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Obtener la ID de la URL
// 		id, err := strconv.Atoi(chi.URLParam(r, "id"))
// 		if err != nil {
// 			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
// 			return
// 		}
// 		// Decodificar el cuerpo de la solicitud
// 		var sectionDoc models.SectionDoc
// 		if err := json.NewDecoder(r.Body).Decode(&sectionDoc); err != nil {
// 			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
// 			return
// 		}

// 		// Convertir SectionDoc a Section
// 		section := ConvertSectionDocToSection(sectionDoc)

// 		// Actualizar la sección
// 		if err := h.sv.PatchSection(id, section); err != nil {
// 			response.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
// 			return
// 		}

// 		// Responder con éxito
// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "section updated successfully",
// 		})
// 	}
// }

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

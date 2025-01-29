package employee

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/employee"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

func NewDefaultHandler(service employee.EmployeeService) *DefaultHandler {
	return &DefaultHandler{sv: service}
}

type DefaultHandler struct {
	sv employee.EmployeeService
}

func (h *DefaultHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		employees, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]dto.EmployeeDoc)
		for key, value := range employees {
			data[key] = dto.EmployeeDoc{
				Id:          value.Id.GetId(),
				CardNumber:  value.CardNumber.GetCardNumber(),
				FirstName:   value.FirstName.GetName(),
				LastName:    value.LastName.GetName(),
				WarehouseId: value.WarehouseId.GetId(),
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *DefaultHandler) GetById() http.HandlerFunc {
	return nil
}

func (h *DefaultHandler) Add() http.HandlerFunc {
	return nil
}

func (h *DefaultHandler) Update() http.HandlerFunc {
	return nil
}

func (h *DefaultHandler) DeleteById() http.HandlerFunc {
	return nil
}

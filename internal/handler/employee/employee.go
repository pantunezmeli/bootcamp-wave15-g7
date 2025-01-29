package employee

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/employee"
)

func NewDefaultHandler(service sv.EmployeeService) *DefaultHandler {
	return &DefaultHandler{sv: service}
}

type DefaultHandler struct {
	sv sv.EmployeeService
}

func (h *DefaultHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - nothing

		// process
		employees, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    employees,
		})
	}
}

func (h *DefaultHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		// process
		employee, err := h.sv.FindById(id)
		if err != nil {
			if errors.Is(err, sv.ErrEmployeeNotFound) {
				response.JSON(w, http.StatusNotFound, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    employee,
		})
	}
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

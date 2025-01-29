package employee

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/employee"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
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
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var employeeData dto.EmployeeDoc
		if err := request.JSON(r, &employeeData); err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		// process
		newEmployee, err := h.sv.New(employeeData)
		if err != nil {
			if errors.Is(err, sv.ErrEmptyField) {
				response.JSON(w, http.StatusUnprocessableEntity, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    newEmployee,
		})
	}
}

func (h *DefaultHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		var employeeData dto.EmployeeDoc
		if err := request.JSON(r, &employeeData); err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		// process
		updatedEmployee, err := h.sv.Edit(id, employeeData)
		if err != nil {
			if errors.Is(err, sv.ErrEmployeeNotFound) {
				response.JSON(w, http.StatusUnprocessableEntity, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    updatedEmployee,
		})
	}
}

func (h *DefaultHandler) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
		}

		// process
		err = h.sv.DeleteById(id)
		if err != nil {
			if errors.Is(err, sv.ErrEmployeeNotFound) {
				response.JSON(w, http.StatusUnprocessableEntity, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "success",
		})
	}
}

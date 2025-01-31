package handler

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
			response.Error(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": employees,
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
				response.Error(w, http.StatusNotFound, "Employee not found")
				return
			}
			response.Error(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": employee,
		})
	}
}

func (h *DefaultHandler) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var employeeData dto.EmployeeDoc
		if err := request.JSON(r, &employeeData); err != nil {
			response.Error(w, http.StatusBadRequest, "Bad request: invalid data")
			return
		}

		// process
		newEmployee, err := h.sv.New(employeeData)
		if err != nil {
			if errors.Is(err, sv.ErrEmptyField) {
				response.Error(w, http.StatusUnprocessableEntity, "Unprocessable entity: empty field/s")
				return
			}
			response.Error(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": newEmployee,
		})
	}
}

func (h *DefaultHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Bad request: invalid id")
		}

		var employeeData dto.EmployeeDoc
		if err := request.JSON(r, &employeeData); err != nil {
			response.Error(w, http.StatusBadRequest, "Bad request: invalid data")
			return
		}

		// process
		updatedEmployee, err := h.sv.Edit(id, employeeData)
		if err != nil {
			if errors.Is(err, sv.ErrEmployeeNotFound) {
				response.Error(w, http.StatusNotFound, "Employee not found")
				return
			}
			response.Error(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": updatedEmployee,
		})
	}
}

func (h *DefaultHandler) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Bad request: invalid id")
		}

		// process
		err = h.sv.DeleteById(id)
		if err != nil {
			if errors.Is(err, sv.ErrEmployeeNotFound) {
				response.Error(w, http.StatusNotFound, "Employee not found")
				return
			}
			response.Error(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, nil)
	}
}

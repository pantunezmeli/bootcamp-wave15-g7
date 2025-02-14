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
	dtoEmployee "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/employee"
)

var (
	ErrEmployeeNotFound          = errors.New("employee not found")
	ErrWarehouseNotFound         = errors.New("warehouse not found")
	ErrCardNumberExists          = errors.New("card number already exists")
	ErrEmptyField                = errors.New("data lacks a required field")
	ErrInboundOrderNeedsEmployee = errors.New("inbound order linked to this employee")
	ErrNotImplemented            = errors.New("not implemented")
)

func NewDefaultHandler(service sv.EmployeeService) *DefaultHandler {
	return &DefaultHandler{sv: service}
}

type DefaultHandler struct {
	sv sv.EmployeeService
}

func (h *DefaultHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - nothing

		// process
		employees, err := h.sv.FindAll()
		if err != nil {
			h.handleError(w, err)
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
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidId.Error())
			return
		}

		// process
		employee, err := h.sv.FindById(id)
		if err != nil {
			h.handleError(w, err)
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": employee,
		})
	}
}

func (h *DefaultHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var employeeData dtoEmployee.EmployeeDoc
		if err := request.JSON(r, &employeeData); err != nil {
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidBody.Error())
			return
		}

		// process
		newEmployee, err := h.sv.New(employeeData)
		if err != nil {
			h.handleError(w, err)
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
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidId.Error())
			return
		}

		var employeeData dtoEmployee.EmployeeDoc
		if err := request.JSON(r, &employeeData); err != nil {
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidBody.Error())
			return
		}

		// process
		updatedEmployee, err := h.sv.Edit(id, employeeData)
		if err != nil {
			h.handleError(w, err)
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": updatedEmployee,
		})
	}
}

func (h *DefaultHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidId.Error())
		}

		// process
		err = h.sv.DeleteById(id)
		if err != nil {
			h.handleError(w, err)
		}

		// response
		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (h *DefaultHandler) ReportInboundOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id := r.URL.Query().Get("id")
		if _, errId := strconv.Atoi(id); errId != nil {
			dto.JSONError(w, http.StatusBadRequest, ErrInvalidId.Error())
			return
		}

		// process
		employee, err := h.sv.ReportInboundOrders(id)
		if err != nil {
			h.handleError(w, err)
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": employee,
		})
	}
}

func (h *DefaultHandler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, sv.ErrEmptyField):
		dto.JSONError(w, http.StatusUnprocessableEntity, ErrEmptyField.Error())
	case errors.Is(err, sv.ErrCardNumberAlreadyExists):
		dto.JSONError(w, http.StatusConflict, ErrCardNumberExists.Error())
	case errors.Is(err, sv.ErrWarehouseNotFound):
		dto.JSONError(w, http.StatusNotFound, ErrWarehouseNotFound.Error())
	case errors.Is(err, sv.ErrEmployeeNotFound):
		dto.JSONError(w, http.StatusNotFound, ErrEmployeeNotFound.Error())
	case errors.Is(err, sv.ErrInboundOrderNeedsEmployee):
		dto.JSONError(w, http.StatusConflict, ErrInboundOrderNeedsEmployee.Error())
	default:
		dto.JSONError(w, http.StatusInternalServerError, ErrInternalServerError.Error())
	}
}

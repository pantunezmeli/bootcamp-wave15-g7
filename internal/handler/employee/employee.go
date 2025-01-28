package employee

import (
	"net/http"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/employee"
)

func NewDefaultHandler(service employee.EmployeeService) *DefaultHandler {
	return &DefaultHandler{sv: service}
}

type DefaultHandler struct {
	sv employee.EmployeeService
}

func (h *DefaultHandler) GetAll() http.HandlerFunc {
	return nil
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

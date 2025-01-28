package employee

import "net/http"

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

type DefaultHandler struct {
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

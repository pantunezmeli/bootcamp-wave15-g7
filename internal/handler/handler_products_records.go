package handler

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product_records"
	"net/http"
)

type HandlerProductRecords struct {
	sv product_records.IProductRecordsService
}

func NewHandlerProductRecords(sv product_records.IProductRecordsService) *HandlerProductRecords {
	return &HandlerProductRecords{sv: sv}
}

func (h HandlerProductRecords) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO implement me
		panic("implement me")
	}
}

func (h HandlerProductRecords) GetRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO implement me
		panic("implement me")
	}
}

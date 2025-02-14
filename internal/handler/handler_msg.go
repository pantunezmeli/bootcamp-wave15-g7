package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
)

// Errors
var (
	MSG_ErrInvalidId              = "invalid Id parameters"
	MSG_ErrIncorrectParameters    = "the fields are empty or incorrect"
	MSG_ErrInternalError          = "internal server error"
	MSG_ErrNotFound               = "element not found"
	MSG_ErrEmptyList              = "the element list is empty"
	MSG_ErrConflict               = "element already exist"
	MSG_ErrStorageOperationFailed = "operation failed in storage"
	MSG_ErrModelInvalid           = "the fields of model is invalid"
	MSG_ErrJsonFormat             = "invalid JSON format"
	MSG_ErrRequest                = "invalid request type"
	MSG_ErrUnprocessable          = "incorrect syntax of the request"
	MSG_ErrOperationDB            = "operation on the database failed"
	MSG_ErrOrderNumberExist       = "cannot repeat the order number in the system"
	MSG_ErrTrackingCodeExist      = "cannot repeat the tracking code in the system"

	// Errors FK
	MSG_ErrBuyerFKNotExist       = "The associated buyer does not exist in the system"
	MSG_ErrCarrierFKNotExist     = "The associated carrier does not exist in the system"
	MSG_ErrOrderStatusFKNotExist = "The associated order status does not exist in the system"
	MSG_ErrWareHouseFKNotExist   = "The associated warehouse does not exist in the system"
	MSG_ErrInvalidIdField        = "the id cannot be zero or empty"
)

// MSG
var (
	MsgSuccess = "success"
	MsgCreated = "element created successfully"
	MsgUpdated = "element uptaded successfully"
)

func jsonResponse(writer http.ResponseWriter, statusCode int, data any) {
	response.JSON(writer, statusCode, map[string]any{
		"data": data,
	})
}

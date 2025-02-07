package handler

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
)

// MSG
var (
	MsgSuccess = "success"
	MsgCreated = "element created successfully"
	MsgUpdated = "element uptaded successfully"
)

package dto

type GenericResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Status  string `json:"status,omitempty"`
}


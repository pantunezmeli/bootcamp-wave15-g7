package dto

import (
	"encoding/json"
	"net/http"

)



func JSONError(w http.ResponseWriter, statusCode int, message string) {
	defaultStatusCode := http.StatusInternalServerError
	if statusCode > 299 && statusCode < 600 {
		defaultStatusCode = statusCode
	}

	body := GenericResponse{
		Status:  http.StatusText(defaultStatusCode),
		Message: message,
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(defaultStatusCode)
	w.Write(bytes)
}

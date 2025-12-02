package controller

import (
	"encoding/json"
	"net/http"
	"subscription-service/internal/entity"
)

func CreateError(code, message string, w http.ResponseWriter) {
	e := entity.ErrorResponse{
		Code:    code,
		Message: message,
	}

	json.NewEncoder(w).Encode(e)
}

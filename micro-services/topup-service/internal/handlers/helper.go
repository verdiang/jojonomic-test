package handlers

import (
	"encoding/json"
	"net/http"
	"topup-service/internal/models"
)

func response(w http.ResponseWriter, code int, err bool, reffID string, data interface{}) {
	response := &models.Response{
		Error:  err,
		ReffID: reffID,
		Data:   data,
	}

	byteResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(byteResponse)
}

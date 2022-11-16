package handlers

import (
	"check-harga-service/internal/models"
	"encoding/json"
	"net/http"
)

func response(w http.ResponseWriter, code int, err bool, data interface{}) {
	response := &models.Response{
		Error: err,
		Data:  data,
	}

	byteResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(byteResponse)
}

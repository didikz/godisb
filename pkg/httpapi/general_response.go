package httpapi

import (
	"encoding/json"
	"net/http"
)

type GeneralResponseError struct {
	Error string `json:"error"`
}

func WriteJson(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

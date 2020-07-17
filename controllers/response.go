package controllers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func newResponse(success bool, message string, data interface{}, err interface{}) response {
	return response{
		success,
		message,
		data,
		err}
}

func newResponseJSON(w http.ResponseWriter, statusCode int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


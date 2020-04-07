package handlers

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Status string `json:"status"`
	Message interface{} `json:"message"`
}
func WriteJSON(w http.ResponseWriter, status int, msg interface{}) {
	res := message{
		"success",
		msg,
	}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

type errorResponse struct {
	Status string `json:"status"`
	Message interface{} `json:"message"`
}

func WriteError(w http.ResponseWriter, status int, msg interface{}) {
	errRes := errorResponse{
		Status:"error",
		Message: msg,
	}
	response, _ := json.Marshal(errRes)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}


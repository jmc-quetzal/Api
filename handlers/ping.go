package handlers

import (
	"net/http"
	"encoding/json"
)

type pong struct {
	Msg	string
}
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(pong{Msg:"Pong"})
}
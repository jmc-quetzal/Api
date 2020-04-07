package handlers

import (
	"net/http"
)

type pong struct {
	Msg string
}

func Ping(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, pong{Msg: "Pong"})
}

func PingAuth(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, pong{Msg: "Auth Pong"})
}
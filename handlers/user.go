package handlers

import (
	"encoding/json"
	"net/http"
	models "github.com/jmc-quetzal/api/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
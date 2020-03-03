package routes

import(
	"github.com/gorilla/mux"
	"github.com/jmc-quetzal/api/handlers"
	
)


func userRoutes(r *mux.Router) {
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("/register",handlers.CreateUserHandler)
}

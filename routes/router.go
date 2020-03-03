package routes

import (
	"github.com/gorilla/mux"
	"github.com/jmc-quetzal/api/handlers"

)


//InitRouter returns a mux Router with routes attached
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping",handlers.Ping)
	userRoutes(r)
	return r
}
package routes

import (
	"github.com/go-chi/chi"
	"github.com/jmc-quetzal/api/config"
	"github.com/jmc-quetzal/api/handlers"
	"github.com/jmc-quetzal/api/postgres"
	"github.com/jmc-quetzal/api/redis"
)

func userRoutes(router *chi.Mux, cfg *config.Config) {
	pgStore := postgres.UserStore{DB: cfg.DB}
	sessionStore := redis.RedisSessions{Pool: cfg.Sessions}
	router.Route("/users", func(router chi.Router) {
		router.Post("/", handlers.CreateUserHandler(pgStore,sessionStore))
		router.Post("/login", handlers.LoginHandler(pgStore,sessionStore))
		router.Delete("/logout",handlers.LogoutHandler(sessionStore))
	})
}

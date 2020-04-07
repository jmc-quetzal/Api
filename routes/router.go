package routes

import (
	"github.com/go-chi/chi"
	"github.com/jmc-quetzal/api/config"
	"github.com/jmc-quetzal/api/handlers"
	"github.com/jmc-quetzal/api/middleware"
	"github.com/jmc-quetzal/api/redis"
)

//InitRouter returns a mux Router with routes attached
func InitRouter(cfg *config.Config) *chi.Mux {
	api := chi.NewRouter()
	api.Get("/ping", handlers.Ping)
	api.Get("/pingAuth",middleware.WithAuth(redis.RedisSessions{cfg.Sessions},handlers.PingAuth))
	userRoutes(api, cfg)
	r := chi.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.CorsMiddleware().Handler)
	r.Mount("/api/", api)
	return r
}

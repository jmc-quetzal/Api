package middleware

import (
	"github.com/rs/cors"
	"os"
)

func CorsMiddleware() *cors.Cors {
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{os.Getenv("ALLOWED_ORIGIN")},
		AllowedMethods:   []string{"GET", "PUT", "HEAD", "POST", "DELETE"},
	})
	return c
}

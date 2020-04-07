package main

import (
	"github.com/jmc-quetzal/api/config"
	"github.com/jmc-quetzal/api/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config, err := config.ApplicationConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := routes.InitRouter(config)
	log.Println("Server starting on port", os.Getenv("PORT"))
	err = http.ListenAndServe(os.Getenv("PORT"), r)

	log.Fatal(err)
}

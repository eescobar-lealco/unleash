package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"leal.co/internal/ports"

	"leal.co/internal/pkg/unleash"
)

func init() {
	err := godotenv.Load() // cargar variables de entorno
	if err != nil {
		log.Printf("Error loading .env, %v", err)
	}

	unleash.Example_unleash()

}

func main() {

	startServer()

}

func startServer() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(apihelpers.JSONContentTypeMiddleware)

	r.Get("/orders", ports.CreateOrder)

	log.Println("starting server")
	http.ListenAndServe(":8080", r)
}

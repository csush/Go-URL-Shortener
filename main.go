package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/csush/Go-URL-Shortener/handlers"
	"github.com/csush/Go-URL-Shortener/models"
	"github.com/csush/Go-URL-Shortener/storage"
)

func main() {
	store := storage.NewStore()
	cg := models.NewCodeGenerator()
	handler := handlers.NewHandler(store, cg)

	router := http.NewServeMux()
	router.HandleFunc("GET /", handler.RedirectURL)
	router.HandleFunc("POST /shorten", handler.ShortenURL)

	server := http.Server{
		Addr:    ":5001",
		Handler: router,
	}

	fmt.Println("Listening on port: 5001")
	log.Fatal(server.ListenAndServe())
}

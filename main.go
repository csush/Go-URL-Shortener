package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ShortenerRequest struct {
	Url string
	Tag string
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})
	router.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		requestData := ShortenerRequest{}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(requestData.Url))
	})

	server := http.Server{
		Addr:    ":5001",
		Handler: router,
	}

	fmt.Println("Listening on port: 5001")
	log.Fatal(server.ListenAndServe())
}

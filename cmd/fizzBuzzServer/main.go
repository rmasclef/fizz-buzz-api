package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/rmasclef/fizz_buzz_api/internal/handler"
)

func main() {
	// define HTTP server
	server := &http.Server{
		Handler: mux.NewRouter().
			HandleFunc("/fizz-buzz", handler.FizzBuzzHandler).
			Methods("POST").
			GetHandler(),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

	// http.HandleFunc("/", fizzBuzzHandler)
	// http.ListenAndServe(":8000", nil)
}

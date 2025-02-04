package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/me/financial/config"
	"github.com/me/financial/handlers"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /person", handlers.CreatePerson)
	mux.HandleFunc("PUT /person", handlers.UpdatePerson)
	mux.HandleFunc("DELETE /person/{id}", handlers.DeletePerson)
	mux.HandleFunc("GET /person/{id}", handlers.FindPersonByID)
	mux.HandleFunc("GET /persons", handlers.FindAllPersons)

	log.Printf("Server running on port %s", config.ServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort()), mux)
}

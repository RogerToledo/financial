package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/me/financial/config"
	"github.com/me/financial/db"
	"github.com/me/financial/handlers"
	"github.com/me/financial/repository"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	rep := repository.NewRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /person", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePerson(rep, w, r)
	})

	mux.HandleFunc("PUT /person/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdatePerson(rep, w, r)
	})

	mux.HandleFunc("DELETE /person/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeletePerson(rep, w, r)
	})

	mux.HandleFunc("GET /person/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindPersonByID(rep, w, r)
	})

	mux.HandleFunc("GET /person/name/{name}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindPersonByName(rep, w, r)
	})

	mux.HandleFunc("GET /persons", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindAllPersons(rep, w, r)
	})

	log.Printf("Server running on port %s", config.ServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort()), mux)
}

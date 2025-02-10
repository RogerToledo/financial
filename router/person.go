package router

import (
	"net/http"

	"github.com/me/financial/repository"

	"github.com/me/financial/handlers"
)

func PersonRoutes(mux *http.ServeMux, rep *repository.Repository) {
	
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

	mux.HandleFunc("GET /persons", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindAllPersons(rep, w, r)
	})
}
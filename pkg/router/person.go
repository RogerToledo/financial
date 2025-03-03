package router

import (
	"net/http"

	"github.com/me/finance/pkg/repository"

	"github.com/me/finance/pkg/controller"
)

func PersonRoutes(mux *http.ServeMux, rep *repository.Repository) {
	c := controller.NewController(rep)
	
	mux.HandleFunc("POST /person", func(w http.ResponseWriter, r *http.Request) {
		c.Person.CreatePerson(rep, w, r)
	})
	
	mux.HandleFunc("PUT /person", func(w http.ResponseWriter, r *http.Request) {
		c.Person.UpdatePerson(rep, w, r)
	})

	mux.HandleFunc("DELETE /person/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Person.DeletePerson(rep, w, r)
	})

	mux.HandleFunc("GET /person/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Person.FindPersonByID(rep, w, r)
	})

	mux.HandleFunc("GET /persons", func(w http.ResponseWriter, r *http.Request) {
		c.Person.FindAllPersons(rep, w, r)
	})
}

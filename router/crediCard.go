package router

import (
	"net/http"

	"github.com/me/financial/handlers"
	"github.com/me/financial/repository"
)

func CreditCardRoutes(mux *http.ServeMux, rep *repository.Repository) {
	mux.HandleFunc("POST /creditCard", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateCreditCard(rep, w, r)
	})

	mux.HandleFunc("PUT /creditCard/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateCreditCard(rep, w, r)
	})

	mux.HandleFunc("DELETE /creditCard/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteCreditCard(rep, w, r)
	})

	mux.HandleFunc("GET /creditCard/{owner}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindCreditCardByOwner(rep, w, r)
	})

	mux.HandleFunc("GET /creditCards", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindAllCreditCard(rep, w, r)
	})
}

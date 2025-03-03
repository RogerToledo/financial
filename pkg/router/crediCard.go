package router

import (
	"net/http"

	"github.com/me/finance/pkg/controller"
	"github.com/me/finance/pkg/repository"
)

func CreditCardRoutes(mux *http.ServeMux, rep *repository.Repository) {
	c := controller.NewController(rep)

	mux.HandleFunc("POST /creditCard", func(w http.ResponseWriter, r *http.Request) {
		c.CreditCard.CreateCreditCard(rep, w, r)
	})

	mux.HandleFunc("PUT /creditCard", func(w http.ResponseWriter, r *http.Request) {
		c.CreditCard.UpdateCreditCard(rep, w, r)
	})

	mux.HandleFunc("DELETE /creditCard/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.CreditCard.DeleteCreditCard(rep, w, r)
	})

	mux.HandleFunc("GET /creditCard/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.CreditCard.FindCreditCardByID(rep, w, r)
	})

	mux.HandleFunc("GET /creditCards", func(w http.ResponseWriter, r *http.Request) {
		c.CreditCard.FindAllCreditCard(rep, w, r)
	})
}

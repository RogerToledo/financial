package router

import (
	"net/http"

	"github.com/me/financial/pkg/repository"

	"github.com/me/financial/pkg/handlers"
)

func PaymentTypeRoutes(mux *http.ServeMux, rep *repository.Repository) {
	
	mux.HandleFunc("POST /paymentType", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePaymentType(rep, w, r)
	})
	
	mux.HandleFunc("PUT /paymentType/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdatePaymentType(rep, w, r)
	})

	mux.HandleFunc("DELETE /paymentType/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeletePaymentType(rep, w, r)
	})

	mux.HandleFunc("GET /paymentType/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindPaymentTypeByID(rep, w, r)
	})

	mux.HandleFunc("GET /paymentType", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindAllPaymentType(rep, w, r)
	})
}
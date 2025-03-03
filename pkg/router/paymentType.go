package router

import (
	"net/http"

	"github.com/me/finance/pkg/repository"

	"github.com/me/finance/pkg/controller"
)

func PaymentTypeRoutes(mux *http.ServeMux, rep *repository.Repository) {
	c := controller.NewController(rep)
	
	mux.HandleFunc("POST /paymentType", func(w http.ResponseWriter, r *http.Request) {
		c.PaymentType.CreatePaymentType(rep, w, r)
	})
	
	mux.HandleFunc("PUT /paymentType", func(w http.ResponseWriter, r *http.Request) {
		c.PaymentType.UpdatePaymentType(rep, w, r)
	})

	mux.HandleFunc("DELETE /paymentType/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.PaymentType.DeletePaymentType(rep, w, r)
	})

	mux.HandleFunc("GET /paymentType/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.PaymentType.FindPaymentTypeByID(rep, w, r)
	})

	mux.HandleFunc("GET /paymentType", func(w http.ResponseWriter, r *http.Request) {
		c.PaymentType.FindAllPaymentType(rep, w, r)
	})
}

package router

import (
	"net/http"

	"github.com/me/finance/pkg/controller"
	"github.com/me/finance/pkg/repository"
)

func InstallmentRoutes(mux *http.ServeMux, rep *repository.Repository) {
	c := controller.NewController(rep)

	mux.HandleFunc("PUT /installment/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Installment.UpdateInstallment(w, r)
	})

	mux.HandleFunc("GET /installment/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Installment.FindInstallmentByPurchaseID(w, r)
	})

	mux.HandleFunc("GET /installment/month/{date}", func(w http.ResponseWriter, r *http.Request) {
		c.Installment.FindInstallmentByMonth(w, r)
	})

	mux.HandleFunc("GET /installment/notPaid", func(w http.ResponseWriter, r *http.Request) {
		c.Installment.FindInstallmentByNotPaid(w, r)
	})
}

package router

import (
	"net/http"

	"github.com/me/financial/pkg/repository"

	"github.com/me/financial/pkg/controller"
)

func PurchaseRoutes(mux *http.ServeMux, rep *repository.Repository) {
	c := controller.NewController(rep)
	
	mux.HandleFunc("POST /purchase", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.CreatePurchase(rep, w, r)
	})
	
	mux.HandleFunc("PUT /purchase", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.UpdatePurchase(rep, w, r)
	})

	mux.HandleFunc("DELETE /purchase/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.DeletePurchase(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindPurchaseByID(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindPurchaseByDate(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/month/{date}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindPurchaseByMonth(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindPurchaseByPerson(rep, w, r)
	})

	mux.HandleFunc("GET /purchase", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindAllPurchases(rep, w, r)
	})
}

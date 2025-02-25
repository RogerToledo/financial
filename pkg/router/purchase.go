package router

import (
	"net/http"

	"github.com/me/financial/pkg/repository"

	"github.com/me/financial/pkg/controller"
)

func PurchaseRoutes(mux *http.ServeMux, rep *repository.Repository) {
	c := controller.NewController(rep)
	
	mux.HandleFunc("POST /purchase", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.Create(rep, w, r)
	})
	
	mux.HandleFunc("PUT /purchase", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.Update(rep, w, r)
	})

	mux.HandleFunc("DELETE /purchase/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.Delete(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindByID(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindByDate(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/month/{date}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindByMonth(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindByPerson(rep, w, r)
	})

	mux.HandleFunc("GET /purchase", func(w http.ResponseWriter, r *http.Request) {
		c.Purchase.FindAll(rep, w, r)
	})
}

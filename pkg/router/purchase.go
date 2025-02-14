package router

import (
	"net/http"

	"github.com/me/financial/pkg/repository"

	"github.com/me/financial/pkg/handlers"
)

func PurchaseRoutes(mux *http.ServeMux, rep *repository.Repository) {
	
	mux.HandleFunc("POST /purchase", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePurchase(rep, w, r)
	})
	
	mux.HandleFunc("PUT /purchase/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdatePurchase(rep, w, r)
	})

	mux.HandleFunc("DELETE /purchase/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeletePurchase(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindPurchaseByID(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/date/{date}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindPurchaseByDate(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/month/{date}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindPurchaseByMonth(rep, w, r)
	})

	mux.HandleFunc("GET /purchase/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindPurchaseByPerson(rep, w, r)
	})

	mux.HandleFunc("GET /purchase", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindAllPurchase(rep, w, r)
	})
}

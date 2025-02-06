package router

import (
	"net/http"

	"github.com/me/financial/repository"

	"github.com/me/financial/handlers"
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

	mux.HandleFunc("GET /purchase", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindAllPurchase(rep, w, r)
	})
}

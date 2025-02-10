package router

import (
	"net/http"

	"github.com/me/financial/pkg/repository"

	"github.com/me/financial/pkg/handlers"
)

func PurchaseTypeRoutes(mux *http.ServeMux, rep *repository.Repository) {
	
	mux.HandleFunc("POST /purchaseType", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePurchaseType(rep, w, r)
	})
	
	mux.HandleFunc("PUT /purchaseType/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdatePurchaseType(rep, w, r)
	})

	mux.HandleFunc("DELETE /purchaseType/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeletePurchaseType(rep, w, r)
	})

	mux.HandleFunc("GET /purchaseType/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindPurchaseTypeByID(rep, w, r)
	})

	mux.HandleFunc("GET /purchaseType", func(w http.ResponseWriter, r *http.Request) {
		handlers.FindAllPurchaseType(rep, w, r)
	})
}

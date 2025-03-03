package router

import (
	"net/http"

	"github.com/me/finance/pkg/repository"

	"github.com/me/finance/pkg/controller"
)

func PurchaseTypeRoutes(mux *http.ServeMux, rep *repository.Repository) {
	c := controller.NewController(rep)
	
	mux.HandleFunc("POST /purchaseType", func(w http.ResponseWriter, r *http.Request) {
		c.PurchaseType.CreatePurchaseType(rep, w, r)
	})
	
	mux.HandleFunc("PUT /purchaseType", func(w http.ResponseWriter, r *http.Request) {
		c.PurchaseType.UpdatePurchaseType(rep, w, r)
	})

	mux.HandleFunc("DELETE /purchaseType/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.PurchaseType.DeletePurchaseType(rep, w, r)
	})

	mux.HandleFunc("GET /purchaseType/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.PurchaseType.FindPurchaseTypeByID(rep, w, r)
	})

	mux.HandleFunc("GET /purchaseType", func(w http.ResponseWriter, r *http.Request) {
		c.PurchaseType.FindAllPurchaseTypes(rep, w, r)
	})
}

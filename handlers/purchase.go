package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/me/financial/model"
	"github.com/me/financial/repository"
)

func CreatePurchase(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var purchase model.Purchase

	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding Purchase: %v", err), http.StatusBadRequest)
		return
	}

	ok, msg := purchase.Validate()
	if !ok {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	id, err := rep.Purchase.Create(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Purchase created with ID: %d", id), http.StatusCreated)
}

func UpdatePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var purchase model.Purchase

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding Purchase: %v", err), http.StatusBadRequest)
		return
	}

	ok, msg := purchase.Validate()
	if !ok {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if err := rep.Purchase.Update(id, purchase); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Purchase updated with ID: %d", id), http.StatusOK)
}

func DeletePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusInternalServerError)
		return		
	}

	err = rep.Purchase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Purchase deleted with ID: %d", id), http.StatusOK)
}

func FindPurchaseByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusInternalServerError)
		return		
	}
	
	purchase, err := rep.Purchase.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, purchase, http.StatusOK)
}

func FindPurchaseByDate(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	date := r.PathValue("date")
	
	purchases, err := rep.Purchase.FindByDate(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchases, http.StatusOK)
}

func FindAllPurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	purchases, err := rep.Purchase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchases, http.StatusOK)
}

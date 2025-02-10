package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

	if err := rep.Purchase.Create(purchase); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Purchase was created with success!"), http.StatusCreated)
}

func UpdatePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var purchase model.Purchase

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
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

	HTTPResponse(w, fmt.Sprintf("Purchase was updated with success!"), http.StatusOK)
}

func DeletePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}

	err = rep.Purchase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Purchase was deleted with success!"), http.StatusOK)
}

func FindPurchaseByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
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

func FindPurchaseByMonth(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	date := r.PathValue("date")
	
	purchases, err := rep.Purchase.FindByMonth(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchases, http.StatusOK)
}

func FindPurchaseByPerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}
	
	purchases, err := rep.Purchase.FindByPerson(id)
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

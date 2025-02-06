package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/me/financial/model"
	"github.com/me/financial/repository"
)

func CreatePurchaseType(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var purchaseType model.PurchaseType

	err := json.NewDecoder(r.Body).Decode(&purchaseType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding Purchase Type: %v", err), http.StatusBadRequest)
		return
	}

	if purchaseType.Name == "" {
		http.Error(w, fmt.Sprint("Name is required"), http.StatusBadRequest)
		return
	}

	id, err := rep.PurchaseType.CreatePurchaseType(purchaseType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Porchase Type created with ID: %d", id), http.StatusCreated)
}

func UpdatePurchaseType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var purchaseType model.PurchaseType

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&purchaseType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding Porchase Type: %v", err), http.StatusBadRequest)
		return
	}

	if purchaseType.Name == "" {
		http.Error(w, fmt.Sprint("Name is required"), http.StatusBadRequest)
		return
	}

	if err := rep.PurchaseType.UpdatePurchaseType(id, purchaseType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Porchase Type updated with ID: %d", id), http.StatusOK)
}

func DeletePurchaseType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusInternalServerError)
		return		
	}

	err = rep.PurchaseType.DeletePurchaseType(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Payment Type deleted with ID: %d", id), http.StatusOK)
}

func FindPurchaseTypeByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusInternalServerError)
		return		
	}
	
	purchaseType, err := rep.PurchaseType.FindPurchaseTypeByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, purchaseType, http.StatusOK)
}

func FindAllPurchaseType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	purchaseTypes, err := rep.PurchaseType.FindAllPurchaseType()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchaseTypes, http.StatusOK)
}

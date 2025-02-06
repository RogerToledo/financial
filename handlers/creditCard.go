package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/me/financial/model"
	"github.com/me/financial/repository"
)

func CreateCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var creditCard model.CreditCard

	if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if creditCard.Owner == "" {
		http.Error(w, fmt.Sprint("Owner is required"), http.StatusBadRequest)
		return
	}

	id, err := rep.CreditCard.CreateCreditCard(creditCard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 

	HTTPResponse(w, fmt.Sprintf("Credit card created with ID: %d", id), http.StatusOK)
}

func UpdateCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var creditCard model.CreditCard

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&creditCard)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if creditCard.Owner == "" {
		http.Error(w, fmt.Sprint("Owner is required"), http.StatusInternalServerError)
		return
	}

	if err = rep.CreditCard.UpdateCreditCard(id, creditCard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Credit card updated with ID: %d", id), http.StatusOK)
}

func DeleteCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusInternalServerError)
		return
	}

	err = rep.CreditCard.DeleteCreditCard(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Credit card deleted with ID: %d", id), http.StatusOK)
}

func FindCreditCardByOwner(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	owner := r.PathValue("owner")
	
	creditCard, err := rep.CreditCard.FindCreditCardByOwner(owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, creditCard, http.StatusOK)
}	

func FindAllCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	creditCard, err := rep.CreditCard.FindAllCreditCards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, creditCard, http.StatusOK)
}
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

	if err := rep.CreditCard.Create(creditCard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 

	HTTPResponse(w, fmt.Sprintf("Credit card created to %s", creditCard.Owner), http.StatusOK)
}

func UpdateCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var creditCard model.CreditCard

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}
	
	if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if creditCard.Owner == "" {
		http.Error(w, fmt.Sprint("Owner is required"), http.StatusInternalServerError)
		return
	}

	if err = rep.CreditCard.Update(id, creditCard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Credit card was updated with success!"), http.StatusOK)
}

func DeleteCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}

	err = rep.CreditCard.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Credit card deleted with sucess!"), http.StatusOK)
}

func FindCreditCardByOwner(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	owner := r.PathValue("owner")
	
	creditCard, err := rep.CreditCard.FindByOwner(owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, creditCard, http.StatusOK)
}	

func FindAllCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	creditCard, err := rep.CreditCard.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, creditCard, http.StatusOK)
}
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/me/financial/model"
	"github.com/me/financial/repository"
)

func CreatePaymentType(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var payment model.PaymentType

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding payment: %v", err), http.StatusBadRequest)
		return
	}

	if payment.Name == "" {
		http.Error(w, fmt.Sprint("Name is required"), http.StatusBadRequest)
		return
	}

	if err := rep.PaymentType.Create(payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Payment Type was created with success!"), http.StatusCreated)
}

func UpdatePaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var payment model.PaymentType

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding Payment Type: %v", err), http.StatusBadRequest)
		return
	}

	if payment.Name == "" {
		http.Error(w, fmt.Sprint("Name is required"), http.StatusBadRequest)
		return
	}

	if err := rep.PaymentType.Update(id, payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Payment Type was updated with success!"), http.StatusOK)
}

func DeletePaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}

	err = rep.PaymentType.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Payment Type was deleted with success!"), http.StatusOK)
}

func FindPaymentTypeByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}
	
	payment, err := rep.PaymentType.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, payment, http.StatusOK)
}

func FindAllPaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	payments, err := rep.PaymentType.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, payments, http.StatusOK)
}

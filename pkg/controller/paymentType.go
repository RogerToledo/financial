package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
	"github.com/me/financial/pkg/usecase"
)

type ControllerPaymentType interface {
	CreatePaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	UpdatePaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	DeletePaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindPaymentTypeByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindAllPaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
}

type controllerPaymentType struct {
	usecase usecase.PaymentTypeUseCase
}

func NewPaymentTypeController(useCase usecase.PaymentTypeUseCase) ControllerPaymentType {
	return &controllerPaymentType{useCase}
}

func (pt *controllerPaymentType) CreatePaymentType(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var paymentType entity.PaymentType

	err := json.NewDecoder(r.Body).Decode(&paymentType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding payment: %v", err), http.StatusBadRequest)
		return
	}

	if err := paymentType.Validate(true); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pt.usecase.CreatePaymentType(paymentType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Payment Type was created with success!"), http.StatusCreated)
}

func (pt *controllerPaymentType) UpdatePaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var paymentType entity.PaymentType

	err := json.NewDecoder(r.Body).Decode(&paymentType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding Payment Type: %v", err), http.StatusBadRequest)
		return
	}

	if err := paymentType.Validate(false); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pt.usecase.UpdatePaymentType(paymentType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Payment Type was updated with success!"), http.StatusOK)
}

func (pt *controllerPaymentType) DeletePaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	if idRequest == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}

	err = pt.usecase.DeletePaymentType(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Payment Type was deleted with success!"), http.StatusOK)
}

func (pt *controllerPaymentType) FindPaymentTypeByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	if idRequest == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}
	
	paymentType, err := pt.usecase.FindPaymentTypeByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, paymentType, http.StatusOK)
}

func (pt *controllerPaymentType) FindAllPaymentType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	paymentTypes, err := pt.usecase.FindAllPaymentTypes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, paymentTypes, http.StatusOK)
}

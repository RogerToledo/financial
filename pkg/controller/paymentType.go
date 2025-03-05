package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/me/finance/pkg/entity"
	"github.com/me/finance/pkg/usecase"
)

type ControllerPaymentType interface {
	CreatePaymentType(w http.ResponseWriter, r *http.Request)
	UpdatePaymentType(w http.ResponseWriter, r *http.Request)
	DeletePaymentType(w http.ResponseWriter, r *http.Request)
	FindPaymentTypeByID(w http.ResponseWriter, r *http.Request)
	FindAllPaymentType(w http.ResponseWriter, r *http.Request)
}

type controllerPaymentType struct {
	usecase usecase.PaymentTypeUseCase
}

func NewPaymentTypeController(useCase usecase.PaymentTypeUseCase) ControllerPaymentType {
	return &controllerPaymentType{useCase}
}

func (pt *controllerPaymentType) CreatePaymentType(w http.ResponseWriter, r *http.Request) {
	var paymentType entity.PaymentType

	err := json.NewDecoder(r.Body).Decode(&paymentType)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding payment: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding payment: %v", err), http.StatusBadRequest)
		return
	}

	if err := paymentType.Validate(true); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pt.usecase.CreatePaymentType(paymentType); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Payment Type was created with success!"), http.StatusCreated)
}

func (pt *controllerPaymentType) UpdatePaymentType(w http.ResponseWriter, r *http.Request) {
	var paymentType entity.PaymentType

	err := json.NewDecoder(r.Body).Decode(&paymentType)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding Payment Type: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding Payment Type: %v", err), http.StatusBadRequest)
		return
	}

	if err := paymentType.Validate(false); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pt.usecase.UpdatePaymentType(paymentType); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Payment Type was updated with success!"), http.StatusOK)
}

func (pt *controllerPaymentType) DeletePaymentType(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pt.usecase.DeletePaymentType(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Payment Type was deleted with success!"), http.StatusOK)
}

func (pt *controllerPaymentType) FindPaymentTypeByID(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")

	id, err := entity.ValidateID(idRequest)
	if err != nil {
		fmt.Sprintf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	paymentType, err := pt.usecase.FindPaymentTypeByID(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, paymentType, http.StatusOK)
}

func (pt *controllerPaymentType) FindAllPaymentType(w http.ResponseWriter, r *http.Request) {
	paymentTypes, err := pt.usecase.FindAllPaymentTypes()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, paymentTypes, http.StatusOK)
}

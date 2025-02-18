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

type ControllerCreditCard interface {
	CreateCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	UpdateCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	DeleteCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindCreditCardByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindAllCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
}

type controllerCreditCard struct{
	useCase usecase.CreditCardUseCase
}

func NewCreditCardController(useCase usecase.CreditCardUseCase) ControllerCreditCard {
	return &controllerCreditCard{useCase}
}

func (c *controllerCreditCard) CreateCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var creditCard entity.CreditCard

	if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if err := creditCard.Validate(true); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.useCase.CreateCreditCard(creditCard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Credit card created to %s", creditCard.Owner), http.StatusOK)
}

func (c *controllerCreditCard) UpdateCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var creditCard entity.CreditCard

	if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if err := creditCard.Validate(false); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.useCase.UpdateCreditCard(creditCard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Credit card was updated with success!"), http.StatusOK)
}

func (c *controllerCreditCard) DeleteCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
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

	if err := c.useCase.DeleteCreditCard(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Credit card deleted with sucess!"), http.StatusOK)
}

func (c *controllerCreditCard) FindCreditCardByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
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

	creditCard, err := c.useCase.FindCreditCardByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	HTTPResponse(w, creditCard, http.StatusOK)
}	

func (c *controllerCreditCard) FindAllCreditCard(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	creditCard, err := c.useCase.FindAllCreditCards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	HTTPResponse(w, creditCard, http.StatusOK)
}

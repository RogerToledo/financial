package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/me/finance/pkg/entity"
	"github.com/me/finance/pkg/usecase"
)

type ControllerCreditCard interface {
	CreateCreditCard(w http.ResponseWriter, r *http.Request)
	UpdateCreditCard(w http.ResponseWriter, r *http.Request)
	DeleteCreditCard(w http.ResponseWriter, r *http.Request)
	FindCreditCardByID(w http.ResponseWriter, r *http.Request)
	FindAllCreditCard(w http.ResponseWriter, r *http.Request)
}

type controllerCreditCard struct{
	useCase usecase.CreditCardUseCase
}

func NewCreditCardController(useCase usecase.CreditCardUseCase) ControllerCreditCard {
	return &controllerCreditCard{useCase}
}

func (c *controllerCreditCard) CreateCreditCard(w http.ResponseWriter, r *http.Request) {
	var creditCard entity.CreditCard

	if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
		slog.Error(fmt.Sprintf("Error decoding credit card: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if err := creditCard.Validate(true); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.useCase.CreateCreditCard(creditCard); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Credit card created to %s", creditCard.Owner), http.StatusOK)
}

func (c *controllerCreditCard) UpdateCreditCard(w http.ResponseWriter, r *http.Request) {
	var creditCard entity.CreditCard

	if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
		slog.Error(fmt.Sprintf("Error decoding credit card: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if err := creditCard.Validate(false); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.useCase.UpdateCreditCard(creditCard); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Credit card was updated with success!"), http.StatusOK)
}

func (c *controllerCreditCard) DeleteCreditCard(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.useCase.DeleteCreditCard(id); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Credit card deleted with sucess!"), http.StatusOK)
}

func (c *controllerCreditCard) FindCreditCardByID(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")

	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	creditCard, err := c.useCase.FindCreditCardByID(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	HTTPResponse(w, creditCard, http.StatusOK)
}	

func (c *controllerCreditCard) FindAllCreditCard(w http.ResponseWriter, r *http.Request) {
	creditCard, err := c.useCase.FindAllCreditCards()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	HTTPResponse(w, creditCard, http.StatusOK)
}

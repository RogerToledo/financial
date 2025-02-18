package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
	"github.com/me/financial/pkg/usecase"
)

type ControllerPurchase interface {
	CreatePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	UpdatePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	DeletePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindPurchaseByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindPurchaseByDate(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindPurchaseByMonth(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindPurchaseByPerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindAllPurchases(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
}

type purchaseController struct {
	useCase usecase.PurchaseUseCase
}

func NewPurchaseController(useCase usecase.PurchaseUseCase) ControllerPurchase {
	return &purchaseController{useCase}
}

func (p *purchaseController) CreatePurchase(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var purchase entity.Purchase

	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding Purchase: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding Purchase: %v", err), http.StatusBadRequest)
		return
	}

	if err := purchase.Validate(); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := rep.Purchase.Create(purchase); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Purchase was created with success!"), http.StatusCreated)
}

func (p *purchaseController) UpdatePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var purchase entity.Purchase

	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding Purchase: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding Purchase: %v", err), http.StatusBadRequest)
		return
	}

	if err := purchase.Validate(); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := rep.Purchase.Update(purchase); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Purchase was updated with success!"), http.StatusOK)
}

func (p *purchaseController) DeletePurchase(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rep.Purchase.Delete(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Purchase was deleted with success!"), http.StatusOK)
}

func (p *purchaseController) FindPurchaseByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		fmt.Sprintf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	purchase, err := rep.Purchase.FindByID(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, purchase, http.StatusOK)
}

func (p *purchaseController) FindPurchaseByDate(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	date := r.PathValue("date")

	if err := entity.ValidateDate(date); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	purchases, err := p.useCase.FindPurchaseByDate(date)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchases, http.StatusOK)
}

func (p *purchaseController) FindPurchaseByMonth(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	date := r.PathValue("date")

	if err := entity.ValidateYearMonth(date); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	purchases, err := p.useCase.FindPurchaseByMonth(date)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchases, http.StatusOK)
}

func (p *purchaseController) FindPurchaseByPerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	purchases, err := p.useCase.FindPurchaseByPerson(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchases, http.StatusOK)
}

func (p *purchaseController) FindAllPurchases(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	purchases, err := p.useCase.FindAllPurchases()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchases, http.StatusOK)
}

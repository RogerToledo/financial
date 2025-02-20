package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/me/financial/pkg/dto"
	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
	"github.com/me/financial/pkg/usecase"
)

type ControllerPurchase interface {
	Create(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	Update(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	Delete(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindByDate(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindByMonth(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindByPerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindAll(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
}

type purchaseController struct {
	useCase usecase.PurchaseUseCase
}

func NewPurchaseController(useCase usecase.PurchaseUseCase) ControllerPurchase {
	return &purchaseController{useCase}
}

func (p *purchaseController) Create(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var (
		purchase entity.Purchase
		purchaseRequest dto.PurchaseRequest
	)	

	err := json.NewDecoder(r.Body).Decode(&purchaseRequest)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding Purchase: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding Purchase: %v", err), http.StatusBadRequest)
		return
	}

	purchase = purchaseRequest.ToEntity()

	if err := purchase.Validate(); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.useCase.CreatePurchase(purchase); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Purchase was created with success!"), http.StatusCreated)
}

func (p *purchaseController) Update(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var (
		purchase entity.Purchase
		purchaseRequest dto.PurchaseRequest
	)

	err := json.NewDecoder(r.Body).Decode(&purchaseRequest)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding Purchase: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding Purchase: %v", err), http.StatusBadRequest)
		return
	}

	purchase = purchaseRequest.ToEntity()

	if err := purchase.Validate(); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.useCase.UpdatePurchase(purchase); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Purchase was updated with success!"), http.StatusOK)
}

func (p *purchaseController) Delete(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = p.useCase.DeletePurchase(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Purchase was deleted with success!"), http.StatusOK)
}

func (p *purchaseController) FindByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	purchase, err := p.useCase.FindPurchaseByID(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, purchase, http.StatusOK)
}

func (p *purchaseController) FindByDate(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
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

func (p *purchaseController) FindByMonth(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
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

func (p *purchaseController) FindByPerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
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

func (p *purchaseController) FindAll(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	purchases, err := p.useCase.FindAllPurchases()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchases, http.StatusOK)
}

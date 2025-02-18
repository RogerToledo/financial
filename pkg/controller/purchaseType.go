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

type ControllerPurchaseType interface {
	CreatePurchaseType(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	UpdatePurchaseType(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	DeletePurchaseType(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindPurchaseTypeByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindAllPurchaseTypes(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
}

type controllerPurchaseType struct{
	usecase usecase.PurchaseTypeUseCase
}

func NewPurchaseTypeController(useCase usecase.PurchaseTypeUseCase) ControllerPurchaseType {
	return &controllerPurchaseType{useCase}
}

func (pt *controllerPurchaseType) CreatePurchaseType(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var purchaseType entity.PurchaseType

	err := json.NewDecoder(r.Body).Decode(&purchaseType)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding Purchase Type: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding Purchase Type: %v", err), http.StatusBadRequest)
		return
	}

	if err := purchaseType.Validate(true); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pt.usecase.CreatePurchaseType(purchaseType); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	HTTPResponse(w, fmt.Sprintf("Purchase Type was created with success!"), http.StatusCreated)
}

func (pt *controllerPurchaseType) UpdatePurchaseType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var purchaseType entity.PurchaseType

	err := json.NewDecoder(r.Body).Decode(&purchaseType)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding Purchase Type: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding Purchase Type: %v", err), http.StatusBadRequest)
		return
	}

	if err := purchaseType.Validate(false); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pt.usecase.UpdatePurchaseType(purchaseType); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Purchase Type was updated with success"), http.StatusOK)
}

func (pt *controllerPurchaseType) DeletePurchaseType(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pt.usecase.DeletePurchaseType(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Purchase Type was deleted with success!"), http.StatusOK)
}

func (pt *controllerPurchaseType) FindPurchaseTypeByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	purchaseType, err := pt.usecase.FindPurchaseTypeByID(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, purchaseType, http.StatusOK)
}

func (pt *controllerPurchaseType) FindAllPurchaseTypes(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	purchaseTypes, err := pt.usecase.FindAllPurchaseTypes()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, purchaseTypes, http.StatusOK)
}

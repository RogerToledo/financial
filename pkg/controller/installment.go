package controller

import (
	"log/slog"
	"net/http"

	"github.com/me/finance/pkg/dto"
	"github.com/me/finance/pkg/entity"
	"github.com/me/finance/pkg/usecase"
)

type ControllerInstallment interface {
	UpdateInstallment(w http.ResponseWriter, r *http.Request)
	FindInstallmentByPurchaseID(w http.ResponseWriter, r *http.Request)
	FindInstallmentByMonth(w http.ResponseWriter, r *http.Request)
	FindInstallmentByNotPaid(w http.ResponseWriter, r *http.Request)
}

type controllerInstallment struct {
	usecase usecase.InstallmentUseCase
}

func NewInstallmentController(useCase usecase.InstallmentUseCase) ControllerInstallment {
	return &controllerInstallment{useCase}
}

func (i *controllerInstallment) UpdateInstallment(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")

	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		HTTPResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := i.usecase.UpdateInstalment(id); err != nil {
		slog.Error(err.Error())
		HTTPResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, "Installment was updated with success!", http.StatusOK)
}

func (i *controllerInstallment) FindInstallmentByPurchaseID(w http.ResponseWriter, r *http.Request) {
	var response dto.InstallmentResponse

	idRequest := r.PathValue("id")

	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		HTTPResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err = i.usecase.FindInstallmentByPurchaseID(id)
	if err != nil {
		slog.Error(err.Error())
		HTTPResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, response, http.StatusOK)
}

func (i *controllerInstallment) FindInstallmentByMonth(w http.ResponseWriter, r *http.Request) {
	month := r.PathValue("date")

	if err := entity.ValidateYearMonth(month); err != nil {
		slog.Error(err.Error())
		HTTPResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	installments, err := i.usecase.FindInstallmentByMonth(month)
	if err != nil {
		slog.Error(err.Error())
		HTTPResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, installments, http.StatusOK)

}

func (i *controllerInstallment) FindInstallmentByNotPaid(w http.ResponseWriter, r *http.Request) {
	installments, err := i.usecase.FindInstallmentByNotPaid()
	if err != nil {
		slog.Error(err.Error())
		HTTPResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	HTTPResponse(w, installments, http.StatusOK)
}

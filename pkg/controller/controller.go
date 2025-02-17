package controller

import (
	"github.com/me/financial/pkg/repository"
	"github.com/me/financial/pkg/usecase"
)

type Controller struct {
	Person       ControllerPerson
	CreditCard   ControllerCreditCard
	PaymentType  ControllerPaymentType
	PurchaseType ControllerPurchaseType
	Purchase     ControllerPurchase
}

func NewController(r *repository.Repository) *Controller {
	p := usecase.NewPersonUseCase(r.Person)
	cc := usecase.NewCreditCardUseCase(r.CreditCard)
	pt := usecase.NewPaymentTypeUseCase(r.PaymentType)
	purt := usecase.NewPurchaseTypeUseCase(r.PurchaseType)
	pur := usecase.NewPurchaseUseCase(r.Purchase)
	

	person       := NewPersonController(p)
	creditCard   := NewCreditCardController(cc)
	paymentType  := NewPaymentTypeController(pt)
	purchaseType := NewPurchaseTypeController(purt)
	purchase     := NewPurchaseController(pur)
	
	return &Controller{
		Person:       person,
		CreditCard:   creditCard,
		PaymentType:  paymentType,
		PurchaseType: purchaseType,
		Purchase:     purchase,
	}
}
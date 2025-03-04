package controller

import (
	"sync"

	"github.com/me/finance/pkg/repository"
	"github.com/me/finance/pkg/usecase"
)

type Controller struct {
	Person       ControllerPerson
	CreditCard   ControllerCreditCard
	PaymentType  ControllerPaymentType
	PurchaseType ControllerPurchaseType
	Purchase     ControllerPurchase
}

var (
	once     sync.Once
	instance *Controller
)

func NewController(r *repository.Repository) *Controller {
	if instance == nil {
		once.Do(func() {
			p := usecase.NewPersonUseCase(r.Person)
			cc := usecase.NewCreditCardUseCase(r.CreditCard)
			pt := usecase.NewPaymentTypeUseCase(r.PaymentType)
			purt := usecase.NewPurchaseTypeUseCase(r.PurchaseType)
			pur := usecase.NewPurchaseUseCase(r.All)
			usecase.NewInstallmentUseCase(r.All)

			person := NewPersonController(p)
			creditCard := NewCreditCardController(cc)
			paymentType := NewPaymentTypeController(pt)
			purchaseType := NewPurchaseTypeController(purt)
			purchase := NewPurchaseController(pur)

			instance = &Controller{
				Person:       person,
				CreditCard:   creditCard,
				PaymentType:  paymentType,
				PurchaseType: purchaseType,
				Purchase:     purchase,
			}
		})
	}	
	
	return instance
}

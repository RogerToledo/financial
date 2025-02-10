package repository

import "database/sql"

type Repository struct {
	Person       *repositoryPerson
	CreditCard   *repositoryCreditCard
	PaymentType  *repositoryPaymentType
	PurchaseType *repositoryPurchaseType
	Purchase     *repositoryPurchase
}

func NewRepository(db *sql.DB) *Repository {
	person       := NewRepositoryPerson(db)
	creditCard   := NewRepositoryCreditCard(db)
	paymentType  := NewRepositoryPaymentType(db)
	purchaseType := NewRepositoryPurchaseType(db)
	purchase     := NewRepositoryPurchase(db)
	
	return &Repository{
		Person:       person,
		CreditCard:   creditCard,
		PaymentType:  paymentType,
		PurchaseType: purchaseType,
		Purchase:     purchase,
	}
}


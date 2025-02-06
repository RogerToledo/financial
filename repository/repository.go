package repository

import "database/sql"

type Repository struct {
	Person      *repositoryPerson
	CreditCard  *repositoryCreditCard
	PaymentType *repositoryPaymentType
}

func NewRepository(db *sql.DB) *Repository {
	person := NewRepositoryPerson(db)
	creditCard := NewRepositoryCreditCard(db)
	paymentType := NewRepositoryPaymentType(db)
	
	return &Repository{
		Person: person,
		CreditCard: creditCard,
		PaymentType: paymentType,
	}
}


package repository

import "database/sql"

type Repository struct {
	Person     *repositoryPerson
	CreditCard *repositoryCreditCard
}

func NewRepository(db *sql.DB) *Repository {
	person := NewRepositoryPerson(db)
	creditCard := NewRepositoryCreditCard(db)
	
	return &Repository{
		Person: person,
		CreditCard: creditCard,
	}
}


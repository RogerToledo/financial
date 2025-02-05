package repository

import "database/sql"

type Repository struct {
	Person *repositoryPerson
}

func NewRepository(db *sql.DB) *Repository {
	person := NewRepositoryPerson(db)
	
	return &Repository{Person: person}
}


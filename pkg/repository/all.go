package repository

import "database/sql"

type RepositoryAll interface {
	All() *Repository
}

type repositoryAll struct {
	db *sql.DB
}

func NewRepositoryAll(db *sql.DB) *repositoryAll {
	return &repositoryAll{db}
}

func (r repositoryAll) All() *Repository {
	return NewRepository(r.db)
}


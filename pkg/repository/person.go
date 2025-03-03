package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/me/finance/pkg/entity"
)

type RepositoryPerson interface {
	Create(p entity.Person) error
	Update(p entity.Person) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (entity.Person, error)
	FindAll() ([]entity.Person, error)
}

type repositoryPerson struct {
	db *sql.DB
}

func NewRepositoryPerson(db *sql.DB) *repositoryPerson {
	return &repositoryPerson{db}
}

func (r repositoryPerson) Create(p entity.Person) error {
	query := `INSERT INTO finance.person (id, name) VALUES ($1, $2)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error trying create uuid: %v", err)
	}

	if _, err = stmt.Exec(id, p.Name); err != nil {
		return fmt.Errorf("error trying insert person: %v", err)
	}

	if err := stmt.Close(); err != nil {
		return fmt.Errorf("error trying close stmt: %v", err)
	}

	return nil
}

func (r repositoryPerson) Update(p entity.Person) error {
	query := `UPDATE finance.person SET name = $1 WHERE id = $2`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	if _, err = stmt.Exec(p.Name, p.ID); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying update person: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist person with this id")
	}

	if err := stmt.Close(); err != nil {
		return fmt.Errorf("error trying close stmt: %v", err)
	}

	return nil
}

func (r repositoryPerson) Delete(id uuid.UUID) error {
	query := `DELETE FROM finance.person WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying delete person: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist person with this id")
	}

	if err := stmt.Close(); err != nil {
		return fmt.Errorf("error trying close stmt: %v", err)
	}

	return nil
}

func (r repositoryPerson) FindByID(id uuid.UUID) (entity.Person, error) {
	query := "SELECT id, name FROM finance.person WHERE id = $1"
	
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return entity.Person{}, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var p entity.Person
	if err = stmt.QueryRow(id).Scan(&p.ID, &p.Name); err != nil && err != sql.ErrNoRows {
		return entity.Person{}, fmt.Errorf("error trying find person: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return entity.Person{}, fmt.Errorf("does not exist person with this id")
	}

	if err := stmt.Close(); err != nil {
		return entity.Person{}, fmt.Errorf("error trying close stmt: %v", err)
	}

	return p, nil
}

func (r repositoryPerson) FindAll() ([]entity.Person, error) {
	query := "SELECT id, name FROM finance.person ORDER BY name"

	rows, err := r.db.Query(query)
	if err != nil {
		return []entity.Person{}, fmt.Errorf("error trying find all persons: %v", err)
	}

	var persons []entity.Person

	for rows.Next() {
		var p entity.Person
		if err = rows.Scan(&p.ID, &p.Name); err != nil && err != sql.ErrNoRows {
			return []entity.Person{}, fmt.Errorf("error trying scan person: %v", err)
		}

		if err != nil && err == sql.ErrNoRows {
			return []entity.Person{}, fmt.Errorf("does not exist person with this name")
		}

		persons = append(persons, p)
	}

	if err := rows.Close(); err != nil {
		return []entity.Person{}, fmt.Errorf("error trying close rows: %v", err)
	}

	return persons, nil
}

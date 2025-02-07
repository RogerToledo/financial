package repository

import (
	"database/sql"
	"fmt"

	"github.com/me/financial/model"
)

type repositoryPerson struct {
	db *sql.DB
}

func NewRepositoryPerson(db *sql.DB) *repositoryPerson {
	return &repositoryPerson{db}
}

func (r repositoryPerson) Create(p model.Person) (int, error) {
	query := `INSERT INTO financial.person (name) VALUES ($1) RETURNING id`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var id int
	if err = stmt.QueryRow(p.Name).Scan(&id); err != nil {
		return 0, fmt.Errorf("error trying insert person: %v", err)
	}

	return id, nil
}

func (r repositoryPerson) Update(id int, p model.Person) error {
	query := `UPDATE financial.person SET name = $1 WHERE id = $2`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	if _, err = stmt.Exec(p.Name, id); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying update person: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist person with this id")
	}

	return nil
}

func (r repositoryPerson) Delete(id int) error {
	query := `DELETE FROM financial.person WHERE id = $1`

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

	return nil
}

func (r repositoryPerson) FindByName(name string) (model.Person, error) {
	query := "SELECT id, name FROM financial.person WHERE name = $1"
	
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return model.Person{}, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var p model.Person
	if err = stmt.QueryRow(name).Scan(&p.ID, &p.Name); err != nil && err != sql.ErrNoRows {
		return model.Person{}, fmt.Errorf("error trying find person: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return model.Person{}, fmt.Errorf("does not exist person with this name")
	}

	return p, nil
}

func (r repositoryPerson) FindAll() ([]model.Person, error) {
	query := "SELECT id, name FROM financial.person"

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.Person{}, fmt.Errorf("error trying find all persons: %v", err)
	}

	var persons []model.Person

	for rows.Next() {
		var p model.Person
		if err = rows.Scan(&p.ID, &p.Name); err != nil && err != sql.ErrNoRows {
			return []model.Person{}, fmt.Errorf("error trying scan person: %v", err)
		}

		if err != nil && err == sql.ErrNoRows {
			return []model.Person{}, fmt.Errorf("does not exist person with this name")
		}

		persons = append(persons, p)
	}

	return persons, nil
}

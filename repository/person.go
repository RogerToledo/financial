package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/me/financial/model"
)

type repositoryPerson struct {
	db *sql.DB
}

func NewRepositoryPerson(db *sql.DB) *repositoryPerson {
	return &repositoryPerson{db}
}

func (r repositoryPerson) CreatePerson(p model.Person) (int, error) {
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

func (r repositoryPerson) UpdatePerson(id int, p model.Person) (int64, error) {
	var rows sql.Result 

	query := `UPDATE financial.person SET name = $1 WHERE id = $2`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error trying prepare statment: %v", err)
	}

	if rows, err = stmt.Exec(p.Name, id); err != nil {
		return 0, fmt.Errorf("error trying update person: %v", err)
	}

	rowAffected, err := rows.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error trying get rows affected: %v", err)
	}

	if rowAffected > 1 {
		log.Printf("error trying update person: more than one person updated: %v", rows)
	}

	return rowAffected, nil
}

func (r repositoryPerson) DeletePerson(id int) (int64, error) {
	query := `DELETE FROM financial.person WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error trying prepare statment: %v", err)
	}

	row, err := stmt.Exec(id)
	if err != nil {
		return 0, fmt.Errorf("error trying delete person: %v", err)
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error trying get rows affected: %v", err)
	}

	if rowAffected > 1 {
		log.Printf("error trying delete person: more than one person deleted: %v", row)
	}

	return rowAffected, nil
}

func (r repositoryPerson) FindPersonByName(name string) (model.Person, error) {
	query := "SELECT id, name FROM financial.person WHERE name = $1"
	
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return model.Person{}, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var p model.Person
	if err = stmt.QueryRow(name).Scan(&p.ID, &p.Name); err != nil {
		return model.Person{}, fmt.Errorf("error trying find person: %v", err)
	}

	return p, nil
}

func (r repositoryPerson) FindAllPersons() ([]model.Person, error) {
	query := "SELECT id, name FROM financial.person"

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.Person{}, fmt.Errorf("error trying find all persons: %v", err)
	}

	var persons []model.Person

	for rows.Next() {
		var p model.Person
		if err = rows.Scan(&p.ID, &p.Name); err != nil {
			return []model.Person{}, fmt.Errorf("error trying scan person: %v", err)
		}

		persons = append(persons, p)
	}

	return persons, nil
}

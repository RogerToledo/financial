package repository

import (
	"database/sql"
	"fmt"

	"github.com/me/financial/model"
)

type repositoryPaymentType struct {
	db *sql.DB
}

func NewRepositoryPaymentType(db *sql.DB) *repositoryPaymentType {
	return &repositoryPaymentType{db}
}

func (r repositoryPaymentType) Create(p model.PaymentType) (int, error) {
	query := `INSERT INTO financial.payment_type (name) VALUES ($1) RETURNING id`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var id int
	if err = stmt.QueryRow(p.Name).Scan(&id); err != nil {
		return 0, fmt.Errorf("error trying insert payment type: %v", err)
	}

	if err := stmt.Close(); err != nil {
		return 0, fmt.Errorf("error trying close stmt: %v", err)
	}

	return id, nil
}

func (r repositoryPaymentType) Update(id int, pt model.PaymentType) error {
	query := `UPDATE financial.payment_type SET name = $1 WHERE id = $2`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	if _, err = stmt.Exec(pt.Name, id); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying update payment type: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist payment type with this id")
	}

	if err := stmt.Close(); err != nil {
		return fmt.Errorf("error trying close stmt: %v", err)
	}

	return nil
}

func (r repositoryPaymentType) Delete(id int) error {
	query := `DELETE FROM financial.payment_type WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying delete payment type: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist payment type with this id")
	}

	if err := stmt.Close(); err != nil {
		return fmt.Errorf("error trying close stmt: %v", err)
	}

	return nil
}

func (r repositoryPaymentType) FindByID(id int) (model.PaymentType, error) {
	query := "SELECT id, name FROM financial.payment_type WHERE id = $1"
	
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return model.PaymentType{}, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var pt model.PaymentType
	if err = stmt.QueryRow(id).Scan(&pt.ID, &pt.Name); err != nil && err != sql.ErrNoRows {
		return model.PaymentType{}, fmt.Errorf("error trying find payment type: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return model.PaymentType{}, fmt.Errorf("does not exist payment type with this id")
	}

	if err := stmt.Close(); err != nil {
		return model.PaymentType{}, fmt.Errorf("error trying close stmt: %v", err)
	}

	return pt, nil
}

func (r repositoryPaymentType) FindAll() ([]model.PaymentType, error) {
	query := "SELECT id, name FROM financial.payment_type ORDER BY name"

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.PaymentType{}, fmt.Errorf("error trying find all payment type: %v", err)
	}

	var payments []model.PaymentType

	for rows.Next() {
		var pt model.PaymentType
		if err = rows.Scan(&pt.ID, &pt.Name); err != nil && err != sql.ErrNoRows {
			return []model.PaymentType{}, fmt.Errorf("error trying scan payment type: %v", err)
		}

		if err != nil && err == sql.ErrNoRows {
			return []model.PaymentType{}, fmt.Errorf("does not exist payment type with this name")
		}

		payments = append(payments, pt)
	}

	if err := rows.Close(); err != nil {
		return []model.PaymentType{}, fmt.Errorf("error trying close rows: %v", err)
	}

	return payments, nil
}

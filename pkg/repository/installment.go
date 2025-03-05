package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/me/finance/pkg/entity"
)

type RepositoryInstallment interface {
	Create(installment entity.Installment) error
	Update(installment entity.Installment) error
	Delete(id uuid.UUID) error
	FindByPurchaseID(id uuid.UUID) ([]entity.Installment, error)
	FindByMonth(month string) ([]entity.Installment, error)
	FindByNotPaid() ([]entity.Installment, error)
}

type repositoryInstallment struct {
  db *sql.DB
}

func NewRepositoryInstallment(db *sql.DB) *repositoryInstallment {
	return &repositoryInstallment{db}
}

func (r *repositoryInstallment) Create(installment entity.Installment) error {
	sql := `INSERT INTO installment (id, description, number, value, month, paid, purchase_id) 
			VALUES 
			($1, $2, $3, $4, $5, $6, $7)`
	
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}

	fmt.Println("purchase_id: ", installment.PurchaseID.String())

	installment.ID = uuid.New()
	
	_, err = stmt.Exec(
		installment.ID,
		installment.Description,
		installment.Number,
		installment.Value,
		installment.Month,
		installment.Paid,
		installment.PurchaseID,
	)
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}
			
	return nil
}

func (r *repositoryInstallment) Update(id uuid.UUID) error {
	sql := `UPDATE installment SET paid = true WHERE id = $1`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}

	return nil
}

func (r *repositoryInstallment) Delete(id uuid.UUID) error {
	sql := `DELETE FROM installment WHERE purchase_id = $1`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}
	
	return nil
}

func (r *repositoryInstallment) FindByPurchaseID(id uuid.UUID) ([]entity.Installment, error) {
	sql := `SELECT id, description, number, value, month, paid, purchase_id
			 FROM installment 
			 WHERE purchase_id = $1`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %v", err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("error executing statement: %v", err)
	}

	var installments []entity.Installment
	for rows.Next() {
		var installment entity.Installment
		err = rows.Scan(
			&installment.ID,
			&installment.Description,
			&installment.Number,
			&installment.Value,
			&installment.Month,
			&installment.Paid,
			&installment.PurchaseID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %v", err)
		}

		installments = append(installments, installment)
	}
		

	return installments, nil
}

func (r *repositoryInstallment) FindByMonth(month string) ([]entity.Installment, error) {
	sql := `SELECT id, description, number, value, month, paid, purchase_id
			 FROM installment 
			 WHERE to_char(month, 'YYYY-MM') = $1`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %v", err)
	}

	rows, err := stmt.Query(month)
	if err != nil {
		return nil, fmt.Errorf("error executing statement: %v", err)
	}

	var installments []entity.Installment
	for rows.Next() {
		var installment entity.Installment
		err = rows.Scan(
			&installment.ID,
			&installment.Description,
			&installment.Number,
			&installment.Value,
			&installment.Month,
			&installment.Paid,
			&installment.PurchaseID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %v", err)
		}

		installments = append(installments, installment)
	}

	return installments, nil
}

func (r *repositoryInstallment) FindByNotPaid() ([]entity.Installment, error) {
	sql := `SELECT id, description, number, value, month, paid, purchase_id 
			FROM installment 
			WHERE paid = false`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %v", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("error executing statement: %v", err)
	}

	var installments []entity.Installment
	for rows.Next() {
		var installment entity.Installment
		err = rows.Scan(
			&installment.ID,
			&installment.Description,
			&installment.Number,
			&installment.Value,
			&installment.Month,
			&installment.Paid,
			&installment.PurchaseID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %v", err)
		}
		
		installments = append(installments, installment)
	}

	return installments, nil
}




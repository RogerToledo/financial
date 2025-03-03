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
	FindByID(id uuid.UUID) (entity.Installment, error)
	FindAll() ([]entity.Installment, error)
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

func (r *repositoryInstallment) Update(entity.Installment) error {
	return nil
}

func (r *repositoryInstallment) Delete(id uuid.UUID) error {
	return nil
}

func (r *repositoryInstallment) FindByID(id uuid.UUID) (entity.Installment, error) {
	return entity.Installment{}, nil
}

func (r *repositoryInstallment) FindAll(entity.Installment) (entity.Installment, error) {
	return entity.Installment{}, nil
}




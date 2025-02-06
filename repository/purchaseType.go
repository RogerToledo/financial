package repository

import (
	"database/sql"
	"fmt"

	"github.com/me/financial/model"
)

type repositoryPurchaseType struct {
	db *sql.DB
}

func NewRepositoryPurchaseType(db *sql.DB) *repositoryPurchaseType {
	return &repositoryPurchaseType{db}
}

func (r repositoryPurchaseType) CreatePurchaseType(p model.PurchaseType) (int, error) {
	query := `INSERT INTO financial.purchase_type (name) VALUES ($1) RETURNING id`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var id int
	if err = stmt.QueryRow(p.Name).Scan(&id); err != nil {
		return 0, fmt.Errorf("error trying insert purchase type type: %v", err)
	}

	return id, nil
}

func (r repositoryPurchaseType) UpdatePurchaseType(id int, pt model.PurchaseType) error {
	query := `UPDATE financial.purchase_type SET name = $1 WHERE id = $2`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	if _, err = stmt.Exec(pt.Name, id); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying update purchase type: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist purchase type with this id")
	}

	return nil
}

func (r repositoryPurchaseType) DeletePurchaseType(id int) error {
	query := `DELETE FROM financial.purchase_type WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying delete purchase type: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist purchase type with this id")
	}

	return nil
}

func (r repositoryPurchaseType) FindPurchaseTypeByID(id int) (model.PurchaseType, error) {
	query := "SELECT id, name FROM financial.purchase_type WHERE id = $1"
	
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return model.PurchaseType{}, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var pt model.PurchaseType
	if err = stmt.QueryRow(id).Scan(&pt.ID, &pt.Name); err != nil && err != sql.ErrNoRows {
		return model.PurchaseType{}, fmt.Errorf("error trying find purchase type: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return model.PurchaseType{}, fmt.Errorf("does not exist purchase type with this id")
	}

	return pt, nil
}

func (r repositoryPurchaseType) FindAllPurchaseType() ([]model.PurchaseType, error) {
	query := "SELECT id, name FROM financial.purchase_type"

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.PurchaseType{}, fmt.Errorf("error trying find all purchase type: %v", err)
	}

	var purchases []model.PurchaseType

	for rows.Next() {
		var pt model.PurchaseType
		if err = rows.Scan(&pt.ID, &pt.Name); err != nil && err != sql.ErrNoRows {
			return []model.PurchaseType{}, fmt.Errorf("error trying scan purchase type: %v", err)
		}

		if err != nil && err == sql.ErrNoRows {
			return []model.PurchaseType{}, fmt.Errorf("does not exist purchase type with this name")
		}

		purchases = append(purchases, pt)
	}

	return purchases, nil
}

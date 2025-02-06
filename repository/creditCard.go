package repository

import (
	"database/sql"
	"fmt"

	"github.com/me/financial/model"
)

type repositoryCreditCard struct {
	db *sql.DB
}

func NewRepositoryCreditCard(db *sql.DB) *repositoryCreditCard {
	return &repositoryCreditCard{db}
}

func (r repositoryCreditCard) CreateCreditCard(cc model.CreditCard) (int, error) {
	query := `INSERT INTO financial.credit_card (owner) VALUES ($1) RETURNING id`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var id int
	if err = stmt.QueryRow(cc.Owner).Scan(&id); err != nil {
		return 0, fmt.Errorf("error trying insert credit card: %v", err)
	}

	return id, nil
}

func (r repositoryCreditCard) UpdateCreditCard(id int, cc model.CreditCard) error {
	query := `UPDATE financial.credit_card SET owner = $1 WHERE id = $2`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	if _, err = stmt.Exec(cc.Owner, id); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying update credit card: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist credit card with this id")
	}

	return nil
}

func (r repositoryCreditCard) DeleteCreditCard(id int) error {
	query := `DELETE FROM financial.credit_card WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying delete credit card: %v", err)
	}
    
	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist credit card with this id")
	}

	return nil
}

func (r repositoryCreditCard) FindCreditCardByOwner(owner string) (model.CreditCard, error) {
	query := "SELECT id, owner FROM financial.credit_card WHERE owner = $1"
	
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return model.CreditCard{}, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var cc model.CreditCard
	if err = stmt.QueryRow(owner).Scan(&cc.ID, &cc.Owner); err != nil  && err != sql.ErrNoRows{
		return model.CreditCard{}, fmt.Errorf("error trying find credit card: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return model.CreditCard{}, fmt.Errorf("does not exist this owner %s", owner)
	}

	return cc, nil
}

func (r repositoryCreditCard) FindAllCreditCards() ([]model.CreditCard, error) {
	query := "SELECT id, owner FROM financial.credit_card"

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.CreditCard{}, fmt.Errorf("error trying find all credit cards: %v", err)
	}

	var creditCards []model.CreditCard

	for rows.Next() {
		var cc model.CreditCard
		if err = rows.Scan(&cc.ID, &cc.Owner); err != nil && err != sql.ErrNoRows {
			return []model.CreditCard{}, fmt.Errorf("error trying scan credit card: %v", err)
		}

		if err != nil && err == sql.ErrNoRows {
			return []model.CreditCard{}, fmt.Errorf("does not exist credit card")
		}

		creditCards = append(creditCards, cc)
	}

	return creditCards, nil
}
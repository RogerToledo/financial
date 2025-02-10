package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/me/financial/model"
)

type repositoryPurchase struct {
	db *sql.DB
}

func NewRepositoryPurchase(db *sql.DB) *repositoryPurchase {
	return &repositoryPurchase{db}
}

func (r repositoryPurchase) Create(p model.Purchase) error {
	query := `INSERT INTO financial.purchase(
		id,
		description, 
		amount, 
		"date", 
		installment_number, 
		installment, 
		place, 
		id_payment_type, 
		id_purchase_type, 
		id_credit_card, 
		id_person
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error trying create uuid: %v", err)
	} 
	
	if _, err = stmt.Exec(
		id,
		p.Description,
		p.Amount,
		p.Date,
		p.InstallmentNumber,
		p.Installment,
		p.Place,
		p.IDPaymentType,
		p.IDPurchaseType,
		p.IDCreditCard,
		p.IDPerson,
	); err != nil {
		return fmt.Errorf("error trying insert purchase type: %v", err)
	}

	if err := stmt.Close(); err != nil {
		return fmt.Errorf("error trying close statment: %v", err)
	}

	return nil
}

func (r repositoryPurchase) Update(id uuid.UUID, p model.Purchase) error {
	query := `UPDATE financial.purchase
		SET description = $1, 
			amount = $2, 
			"date" = $3, 
			installment_number = $4, 
			installment = $5, 
			place = $6, 
			id_payment_type = $7, 
			id_purchase_type = $8, 
			id_credit_card = $9, 
			id_person = $10
		WHERE id=$11;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	if _, err = stmt.Exec(
		p.Description,
		p.Amount,
		p.Date,
		p.InstallmentNumber,
		p.Installment,
		p.Place,
		p.IDPaymentType,
		p.IDPurchaseType,
		p.IDCreditCard,
		p.IDPerson,
		id); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying update purchase: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist purchase with this id")
	}

	if err := stmt.Close(); err != nil {
		return fmt.Errorf("error trying close statment: %v", err)
	}

	return nil
}

func (r repositoryPurchase) Delete(id uuid.UUID) error {
	query := `DELETE FROM financial.purchase WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error trying prepare statment: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error trying delete purchase: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return fmt.Errorf("does not exist purchase with this id")
	}

	if err := stmt.Close(); err != nil {
		return fmt.Errorf("error trying close statment: %v", err)
	}

	return nil
}

func (r repositoryPurchase) FindByID(id uuid.UUID) (model.PurchaseResponse, error) {
	query := `SELECT 
				p.id, 
				p.description, 
				p.amount, 
				p."date", 
				p.installment_number, 
				p.installment, 
				p.place, 
				pt."name",
				purt."name", 
				cc."owner", 
				per."name" 
			FROM financial.purchase p
			inner join financial.payment_type pt 
				on p.id_payment_type = pt.id 
			inner join financial.purchase_type purt	
				on p.id_purchase_type = purt.id 
			inner join financial.credit_card cc	
				on p.id_credit_card = cc.id
			inner join financial.person per	
				on p.id_person = per.id 	
			WHERE p.id = $1;`
	
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return model.PurchaseResponse{}, fmt.Errorf("error trying prepare statment: %v", err)
	}

	var pt model.PurchaseResponse
	if err = stmt.QueryRow(id).Scan(
		&pt.ID, 
		&pt.Description,
		&pt.Amount,
		&pt.Date,
		&pt.InstallmentNumber,
		&pt.Installment,
		&pt.Place,
		&pt.PaymentType,
		&pt.PurchaseType,
		&pt.CreditCard,
		&pt.Person,
		); err != nil && err != sql.ErrNoRows {
		return model.PurchaseResponse{}, fmt.Errorf("error trying find purchase: %v", err)
	}

	if err != nil && err == sql.ErrNoRows {
		return model.PurchaseResponse{}, fmt.Errorf("does not exist purchase with this id")
	}

	if err := stmt.Close(); err != nil {
		return model.PurchaseResponse{}, fmt.Errorf("error trying close statment: %v", err)
	}

	return pt, nil
}

func (r repositoryPurchase) FindByDate(date string) (model.PurchaseResponseDate, error) {
	var (
		total float64
		purchases []model.PurchaseResponse
		responses model.PurchaseResponseDate
	)	

	query := `SELECT 
				p.id, 
				p.description, 
				p.amount, 
				p."date", 
				p.installment_number, 
				p.installment, 
				p.place, 
				pt."name",
				purt."name", 
				cc."owner", 
				per."name"
			FROM financial.purchase p
			inner join financial.payment_type pt 
				on p.id_payment_type = pt.id 
			inner join financial.purchase_type purt	
				on p.id_purchase_type = purt.id 
			inner join financial.credit_card cc	
				on p.id_credit_card = cc.id
			inner join financial.person per	
				on p.id_person = per.id 
			WHERE "date" = $1;`

	rows, err := r.db.Query(query, date)
	if err != nil {
		return model.PurchaseResponseDate{}, fmt.Errorf("error trying find purchase by date: %v", err)
	}

	for rows.Next() {
		var p model.PurchaseResponse
		if err = rows.Scan(
			&p.ID,
			&p.Description,
			&p.Amount,
			&p.Date,
			&p.InstallmentNumber,
			&p.Installment,
			&p.Place,
			&p.PaymentType,
			&p.PurchaseType,
			&p.CreditCard,
			&p.Person,
		    ); err != nil && err != sql.ErrNoRows {
			return model.PurchaseResponseDate{}, fmt.Errorf("error trying scan purchase: %v", err)
		}

		if err != nil && err == sql.ErrNoRows {
			return model.PurchaseResponseDate{}, fmt.Errorf("does not exist purchase  with this date")
		}

		purchases = append(purchases, p)
	}

	for _, pur := range purchases {
		total += pur.Amount
	}

	responses = model.PurchaseResponseDate{
		Responses: purchases,
		Total: total,
	}

	if err := rows.Close(); err != nil {
		return model.PurchaseResponseDate{}, fmt.Errorf("error trying close rows: %v", err)
	}

	return responses, nil
}

func (r repositoryPurchase) FindByMonth(date string) (model.PurchaseResponseDate, error){
	var (
		total float64
		purchases []model.PurchaseResponse
		responses model.PurchaseResponseDate
	)	

	query := `SELECT 
				p.id, 
				p.description, 
				p.amount, 
				p."date", 
				p.installment_number, 
				p.installment, 
				p.place, 
				pt."name",
				purt."name", 
				cc."owner", 
				per."name"
			FROM financial.purchase p
			inner join financial.payment_type pt 
				on p.id_payment_type = pt.id 
			inner join financial.purchase_type purt	
				on p.id_purchase_type = purt.id 
			inner join financial.credit_card cc	
				on p.id_credit_card = cc.id
			inner join financial.person per	
				on p.id_person = per.id 
			WHERE to_char(p."date", 'YYYY-MM') = $1
			ORDER BY p."date";`

	rows, err := r.db.Query(query, date)
	if err != nil {
		return model.PurchaseResponseDate{}, fmt.Errorf("error trying find purchase by date: %v", err)
	}

	for rows.Next() {
		var p model.PurchaseResponse
		if err = rows.Scan(
			&p.ID,
			&p.Description,
			&p.Amount,
			&p.Date,
			&p.InstallmentNumber,
			&p.Installment,
			&p.Place,
			&p.PaymentType,
			&p.PurchaseType,
			&p.CreditCard,
			&p.Person,
		    ); err != nil && err != sql.ErrNoRows {
			return model.PurchaseResponseDate{}, fmt.Errorf("error trying scan purchase: %v", err)
		}

		if err != nil && err == sql.ErrNoRows {
			return model.PurchaseResponseDate{}, fmt.Errorf("does not exist purchase  with this date")
		}

		purchases = append(purchases, p)
	}

	for _, pur := range purchases {
		total += pur.Amount
	}

	responses = model.PurchaseResponseDate{
		Responses: purchases,
		Total: total,
	}

	if err := rows.Close(); err != nil {
		return model.PurchaseResponseDate{}, fmt.Errorf("error trying close rows: %v", err)
	}

	return responses, nil
}

func (r repositoryPurchase) FindAll() ([]model.Purchase, error) {
	query := `SELECT 
				id, 
				description, 
				amount, 
				"date", 
				installment_number, 
				installment, 
				place, 
				id_payment_type, 
				id_purchase_type, 
				id_credit_card, 
				id_person
			FROM financial.purchase
			ORDER BY "date";`

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.Purchase{}, fmt.Errorf("error trying find all purchase: %v", err)
	}

	var purchases []model.Purchase

	for rows.Next() {
		var p model.Purchase
		if err = rows.Scan(
			&p.ID,
			&p.Description,
			&p.Amount,
			&p.Date,
			&p.InstallmentNumber,
			&p.Installment,
			&p.Place,
			&p.IDPaymentType,
			&p.IDPurchaseType,
			&p.IDCreditCard,
			&p.IDPerson,
		    ); err != nil && err != sql.ErrNoRows {
			return []model.Purchase{}, fmt.Errorf("error trying scan purchase: %v", err)
		}

		if err != nil && err == sql.ErrNoRows {
			return []model.Purchase{}, fmt.Errorf("does not exist purchase")
		}

		purchases = append(purchases, p)
	}

	if err := rows.Close(); err != nil {
		return []model.Purchase{}, fmt.Errorf("error trying close rows: %v", err)
	}

	return purchases, nil
}

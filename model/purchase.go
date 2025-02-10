package model

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Purchase struct {
	ID                uuid.UUID `json:"id"`	
	Description       string  `json:"description"`
	Type 	          string  `json:"type"`
	Amount            float64 `json:"amount"`
	Date              string  `json:"date"` 
	InstallmentNumber int     `json:"installment_number"`
	Installment       float64 `json:"installment"`
	Place	          string  `json:"place"`
	IDPaymentType     uuid.UUID `json:"id_payment_type"`
	IDCreditCard	  uuid.UUID `json:"id_credit_card"`
	IDPurchaseType    uuid.UUID `json:"id_purchase_type"`
	IDPerson	      uuid.UUID `json:"id_person"`
}

type PurchaseResponse struct {
	ID                uuid.UUID `json:"id"`
	Description       string  `json:"description"`
	Amount            float64 `json:"amount"`
	Date              string  `json:"date"` 
	InstallmentNumber int     `json:"installment_number"`
	Installment       float64 `json:"installment"`
	Place	          string  `json:"place"`
	PaymentType       string  `json:"payment_type"`
	CreditCard	      string  `json:"credit_card"`
	PurchaseType      string  `json:"purchase_type"`
	Person	          string  `json:"person"`
}

type PurchaseResponseDate struct {
	Responses []PurchaseResponse `json:"responses"`
	Total     float64            `json:"total"`
}

func (p *Purchase) Validate() (bool, string) {
	var invalidFields []string

	if p.Amount <= 0{
		invalidFields = append(invalidFields, "Amount")
	}

	if p.Date == "" {
		invalidFields = append(invalidFields, "Data")
	}

	if p.IDPaymentType == uuid.Nil {
		invalidFields = append(invalidFields, "ID of Payment Type")
	}

	if p.IDCreditCard == uuid.Nil {
		invalidFields = append(invalidFields, "ID of Credit Card")
	}

	if p.IDPurchaseType == uuid.Nil {
		invalidFields = append(invalidFields, "ID of Purchase Type")
	}

	if p.IDPerson == uuid.Nil {
		invalidFields = append(invalidFields, "ID of Person")
	}

	if len(invalidFields) > 0 {
		fields := strings.Join(invalidFields, ", ")

		return false, fmt.Sprintf("The field(s) %s are required", fields)
	}

	return true, ""
}

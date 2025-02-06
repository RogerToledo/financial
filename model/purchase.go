package model

import (
	"fmt"
	"strings"
)

type Purchase struct {
	ID                int     `json:"id"`	
	Description       string  `json:"description"`
	Type 	          string  `json:"type"`
	Amount            float64 `json:"amount"`
	Date              string  `json:"date"` 
	InstallmentNumber int     `json:"installment_number"`
	Installment       float64 `json:"installment"`
	Place	          string  `json:"place"`
	IDPaymentType     int     `json:"id_payment_type"`
	IDCreditCard	  int     `json:"id_credit_card"`
	IDPurchaseType    int     `json:"id_purchase_type"`
	IDPerson	      int     `json:"id_person"`
}

type PurchaseResponse struct {
	ID                int     `json:"id"`	
	Description       string  `json:"description"`
	Type 	          string  `json:"type"`
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

func (p *Purchase) Validate() (bool, string) {
	var invalidFields []string

	if p.Amount <= 0{
		invalidFields = append(invalidFields, "Amount")
	}

	if p.Date == "" {
		invalidFields = append(invalidFields, "Data")
	}

	if p.IDPaymentType <= 0 {
		invalidFields = append(invalidFields, "ID of Payment Type")
	}

	if p.IDCreditCard <= 0 {
		invalidFields = append(invalidFields, "ID of Credit Card")
	}

	if p.IDPurchaseType <= 0 {
		invalidFields = append(invalidFields, "ID of Purchase Type")
	}

	if p.IDPerson <= 0 {
		invalidFields = append(invalidFields, "ID of Person")
	}

	if len(invalidFields) > 0 {
		fields := strings.Join(invalidFields, ", ")

		return false, fmt.Sprintf("The field(s) %s are required", fields)
	}

	return true, ""
}

package entity

import (
	"fmt"
	"strings"
	"time"

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

type PurchaseResponseTotal struct {
	Responses []PurchaseResponse `json:"responses"`
	Quantity  int                `json:"quantity"`
	Total     float64            `json:"total"`
}

func (p *Purchase) Validate() error {
	var (
		invalidFields []string
		msg string
	)	

	if p.Amount <= 0{
		invalidFields = append(invalidFields, "Amount")
	}

	if p.Date == "" {
		invalidFields = append(invalidFields, "Data")
	}

	if err := ValidateDate(p.Date); err != nil {
		invalidFields = append(invalidFields, "Valid Data")
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

		if len(invalidFields) == 1 {
			msg = fmt.Sprintf("The field %s is required", fields)
		} else {
			msg = fmt.Sprintf("The fields %s are required", fields)
		}
	}

	field := validateInstallment(p)
	if field == "" && msg != "" {
		return fmt.Errorf("%s", msg)
	}

	if field != "" && msg == "" {
		return fmt.Errorf("%s", field)
	}

	if field != "" && msg != "" {
		return fmt.Errorf("%s \n%s", msg, field)
	}

	return nil
}
func ValidateDate(date string) error {
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return fmt.Errorf("The date is invalid")
	}

	return nil
}

func ValidateYearMonth(date string) error {
	if _, err := time.Parse("2006-01", date); err != nil {
		return fmt.Errorf("The date is invalid")
	}

	return nil
}

func validateInstallment(p *Purchase) string {
	if p.Installment == 0 && p.InstallmentNumber == 0 {
		return ""
	}

	if p.Installment != 0 && p.InstallmentNumber == 0 {
		return "If Installments was greater than zero, Number of installments is required"
	}

	if p.InstallmentNumber != 0 && p.Installment == 0 {
		return "If Number of installments was greater than zero, Installments is required"
	}

	installmentTotal := p.Installment * float64(p.InstallmentNumber)

	if installmentTotal != p.Amount {
		return "The amount of installment is different from the amount of purchase"
	}

	return ""
}

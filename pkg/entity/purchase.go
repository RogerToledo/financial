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
	Installment       Installment `json:"installment_number"`
	Paid			  bool    `json:"paid"`
	Place	          string  `json:"place"`
	IDPaymentType     uuid.UUID `json:"id_payment_type"`
	IDCreditCard	  uuid.UUID `json:"id_credit_card"`
	IDPurchaseType    uuid.UUID `json:"id_purchase_type"`
	IDPerson	      uuid.UUID `json:"id_person"`
}

func (p *Purchase) Validate() error {
	var invalidFields []string

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
			return fmt.Errorf("The field %s is required", fields)
		} else {
			return fmt.Errorf("The fields %s are required", fields)
		}
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
	if _, err := time.Parse("01/2006", date); err != nil {
		return fmt.Errorf("The date is invalid")
	}

	return nil
}

package entity

import "github.com/google/uuid"

type Installment struct {
	ID          uuid.UUID `son:"id"`
	PurchaseID  uuid.UUID `json:"purchase_id"`
	Description string    `json:"description"`
	Number      int     `json:"number"`
	Value       float64 `json:"value"`
	Month       string  `json:"month"`
	Paid 	    bool    `json:"paid"`
}

type InstallmentRequest struct {
	ID     uuid.UUID `json:"id"`
	Number int       `json:"number"`
	Value  float64 `json:"value"`
	Month  int     `json:"month"`
	Paid   bool    `json:"paid"`
}
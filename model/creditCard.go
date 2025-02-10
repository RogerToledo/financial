package model

import "github.com/google/uuid"

type CreditCard struct {
	ID    uuid.UUID `json:"id"`
	Owner string    `json:"owner"`
}

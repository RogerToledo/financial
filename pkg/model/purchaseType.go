package model

import "github.com/google/uuid"

type PurchaseType struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

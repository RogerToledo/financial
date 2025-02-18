package entity

import (
	"fmt"
	"github.com/google/uuid"
)

func ValidateID(idRequest string) (uuid.UUID, error) {
	id, err := uuid.Parse(idRequest)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error converting ID to UUID: %v", err)
	}

	return id, nil
}
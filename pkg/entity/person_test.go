package entity

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestValidate (t *testing.T) {
	testCases := []struct{
		description string
		person      Person
		removeID    bool
		expected    error
	}{
		{
			description: "Valid person with removeID false",
			person: Person{
				ID:   uuid.New(),
				Name: "John Doe",
			},
			removeID: false,
			expected: nil,
		},
		{
			description: "Empty name with removeID false",
			person: Person{
				ID:   uuid.New(),
				Name: "",
			},
			removeID: false,
			expected: fmt.Errorf("The field name is required"),
		},
		{
			description: "Empty id with removeID false",
			person: Person{
				ID:   uuid.Nil,
				Name: "John Doe",
			},
			removeID: false,
			expected: fmt.Errorf("The field ID is required"),
		},
		{
			description: "Empty all fields with removeID false",
			person: Person{
				ID:   uuid.Nil,
				Name: "",
			},
			removeID: false,
			expected: fmt.Errorf("The fields ID, name are required"),
		},
		{
			description: "Valid person with removeID true",
			person: Person{
				ID:   uuid.New(),
				Name: "John Doe",
			},
			removeID: true,
			expected: nil,
		},
		{
			description: "Empty name with removeID true",
			person: Person{
				ID:   uuid.New(),
				Name: "",
			},
			removeID: true,
			expected: fmt.Errorf("The field name is required"),
		},
		{
			description: "Empty id with removeID true",
			person: Person{
				ID:   uuid.Nil,
				Name: "John Doe",
			},
			removeID: true,
			expected: nil,
		},
		{
			description: "Empty all fields with removeID true",
			person: Person{
				ID:   uuid.Nil,
				Name: "",
			},
			removeID: true,
			expected: fmt.Errorf("The field name is required"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := tc.person.Validate(tc.removeID)

			if result == nil && tc.expected != nil || result != nil && tc.expected == nil {
                t.Errorf("Expected %v, got %v", tc.expected, result)
            } else if result != nil && tc.expected != nil && result.Error() != tc.expected.Error() {
                t.Errorf("Expected %v, got %v", tc.expected.Error(), result.Error())
            }
		})
	}
}

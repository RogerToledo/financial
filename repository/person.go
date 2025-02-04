package repository

import "github.com/me/financial/model"

func CreatePerson(p model.Person) (int, error) {
	return 1, nil
}

func UpdatePerson(id int, p model.Person) (int, error) {
	return 1, nil
}

func DeletePerson(id int) (int, error) {
	return 1, nil
}

func FindPersonByID(id int) (model.Person, error) {
	return model.Person{
		ID: 1,
		Name: "John Doe",
	}, nil
}

func FindAllPersons() ([]model.Person, error) {
	return []model.Person{
		{
			ID: 1,
			Name: "John Doe",
		},
		{
			ID: 2,
			Name: "Jane Doe",
		},
	}, nil
}

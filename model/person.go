package model

type Person struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
}

func (p Person) Create() {
	println("Creating a person")
}

func (p Person) Update() {
	println("Updating a person")
}

func (p Person) Delete() {
	println("Deleting a person")
}

func (p Person) FindByID() {
	println("Finding a person by ID")
}

func (p Person) FindAll() {
	println("Finding all persons")
}

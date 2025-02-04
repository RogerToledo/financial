package model

type PaymentType struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
}

func (p PaymentType) Create() {
	println("Creating a payment type")
}

func (p PaymentType) Update() {
	println("Updating a payment type")
}

func (p PaymentType) Delete() {
	println("Deleting a payment type")
}

func (p PaymentType) FindByID() {
	println("Finding a payment type by ID")
}

func (p PaymentType) FindAll() {
	println("Finding all payment types")
}

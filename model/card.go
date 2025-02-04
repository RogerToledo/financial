package model

type Card struct {
	ID          int       `json:"id"`
	Owner       string    `json:"name"`
}

func (c Card) Create() {
	println("Creating a card")
}

func (c Card) Update() {
	println("Updating a card")
}

func (c Card) Delete() {
	println("Deleting a card")
}

func (c Card) FindByID() {
	println("Finding a card by ID")
}

func (c Card) FindAll() {
	println("Finding all cards")
}

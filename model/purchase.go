package model

type Purchase struct {
	ID          int       `json:"id"`	
	Description string    `json:"description"`
	Type 	    string    `json:"type"`
	Amount      float64   `json:"amount"`
	Date        string    `json:"date"` 
	InstallmentNumber int `json:"installment_number"`
	Installment float64   `json:"installment"`
	Place	    string    `json:"place"`
	IDPayment   int       `json:"id_payment"`
	IDCard	    int       `json:"id_card"`
	IDType	    int       `json:"id_type"`
	IDPerson	int       `json:"id_person"`
}

func (p Purchase) Create() {
	println("Creating a purchase")
}

func (p Purchase) Update() {
	println("Updating a purchase")
}

func (p Purchase) Delete() {
	println("Deleting a purchase")
}

func (p Purchase) FindByID() {
	println("Finding a purchase by ID")
}

func (p Purchase) FindAll() {
	println("Finding all purchases")
}

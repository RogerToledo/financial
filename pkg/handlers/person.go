package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/me/financial/pkg/model"
	"github.com/me/financial/pkg/repository"
)

func CreatePerson(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var person model.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding person: %v", err), http.StatusBadRequest)
		return
	}

	if person.Name == "" {
		http.Error(w, fmt.Sprint("Person is required"), http.StatusBadRequest)
		return
	}

	if err := rep.Person.Create(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Person was created with success!"), http.StatusCreated)
}

func UpdatePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var person model.Person

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if person.Name == "" {
		http.Error(w, fmt.Sprint("Name is required"), http.StatusBadRequest)
		return
	}

	if err := rep.Person.Update(id, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Person was updated with success!"), http.StatusOK)
}

func DeletePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusBadRequest)
		return
	}

	err = rep.Person.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Person was deleted with success!"), http.StatusOK)
}

func FindPersonByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusBadRequest)
		return
	}
	
	person, err := rep.Person.FindByName(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, person, http.StatusOK)
}

func FindAllPersons(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	persons, err := rep.Person.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, persons, http.StatusOK)
}

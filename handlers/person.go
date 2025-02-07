package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/me/financial/model"
	"github.com/me/financial/repository"
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

	id, err := rep.Person.Create(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Person created with ID: %d", id), http.StatusCreated)
}

func UpdatePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var person model.Person

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusBadRequest)
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

	HTTPResponse(w, fmt.Sprintf("Person updated with ID: %d", id), http.StatusOK)
}

func DeletePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to integer: %v", err), http.StatusInternalServerError)
		return		
	}

	err = rep.Person.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Person deleted with ID: %d", id), http.StatusOK)
}

func FindPersonByName(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	
	person, err := rep.Person.FindByName(name)
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

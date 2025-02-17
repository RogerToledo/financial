package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
	"github.com/me/financial/pkg/usecase"
)

type ControllerPerson interface {
	CreatePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	UpdatePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	DeletePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindPersonByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
	FindAllPersons(rep *repository.Repository, w http.ResponseWriter, r *http.Request)
}

type controllerPerson struct{
	useCase usecase.PersonUseCase
}

func NewPersonController(useCase usecase.PersonUseCase) ControllerPerson {
	return &controllerPerson{useCase}
}

func (p *controllerPerson) CreatePerson(rep *repository.Repository , w http.ResponseWriter, r *http.Request) {
	var person entity.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding person: %v", err), http.StatusBadRequest)
		return
	}

	if err := person.Validate(true); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.useCase.CreatePerson(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Person was created with success!"), http.StatusCreated)
}

func (p *controllerPerson) UpdatePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	var person entity.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if err := person.Validate(false); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.useCase.UpdatePerson(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Person was updated with success!"), http.StatusOK)
}

func (p *controllerPerson) DeletePerson(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	if idRequest == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusBadRequest)
		return
	}

	if err := p.useCase.DeletePerson(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Person was deleted with success!"), http.StatusOK)
}

func (p *controllerPerson) FindPersonByID(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	if idRequest == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
	}

	id, err := uuid.Parse(idRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting ID to UUID: %v", err), http.StatusBadRequest)
		return
	}
	
	person, err := p.useCase.FindPersonByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, person, http.StatusOK)
}

func (p *controllerPerson) FindAllPersons(rep *repository.Repository, w http.ResponseWriter, r *http.Request) {
	persons, err := p.useCase.FindAllPersons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, persons, http.StatusOK)
}

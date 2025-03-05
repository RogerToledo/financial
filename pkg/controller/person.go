package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/me/finance/pkg/entity"
	"github.com/me/finance/pkg/usecase"
)

type ControllerPerson interface {
	CreatePerson(w http.ResponseWriter, r *http.Request)
	UpdatePerson(w http.ResponseWriter, r *http.Request)
	DeletePerson(w http.ResponseWriter, r *http.Request)
	FindPersonByID(w http.ResponseWriter, r *http.Request)
	FindAllPersons(w http.ResponseWriter, r *http.Request)
}

type controllerPerson struct{
	useCase usecase.PersonUseCase
}

func NewPersonController(useCase usecase.PersonUseCase) ControllerPerson {
	return &controllerPerson{useCase}
}

func (p *controllerPerson) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person entity.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding person: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding person: %v", err), http.StatusBadRequest)
		return
	}

	if err := person.Validate(true); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.useCase.CreatePerson(person); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Person was created with success!"), http.StatusCreated)
}

func (p *controllerPerson) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var person entity.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		slog.Error(fmt.Sprintf("Error decoding credit card: %v", err))
		http.Error(w, fmt.Sprintf("Error decoding credit card: %v", err), http.StatusBadRequest)
		return
	}

	if err := person.Validate(false); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.useCase.UpdatePerson(person); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprintf("Person was updated with success!"), http.StatusOK)
}

func (p *controllerPerson) DeletePerson(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")

	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := p.useCase.DeletePerson(id); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, fmt.Sprint("Person was deleted with success!"), http.StatusOK)
}

func (p *controllerPerson) FindPersonByID(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")
	
	id, err := entity.ValidateID(idRequest)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	person, err := p.useCase.FindPersonByID(id)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HTTPResponse(w, person, http.StatusOK)
}

func (p *controllerPerson) FindAllPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := p.useCase.FindAllPersons()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HTTPResponse(w, persons, http.StatusOK)
}

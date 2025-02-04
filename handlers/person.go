package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/me/financial/model"
	"github.com/me/financial/repository"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var (
		person model.Person
		resp map[string]any
	)	

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Printf("Error decoding person: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.CreatePerson(person)

	if err != nil {
		resp = map[string]any{
			"StatusCode": http.StatusInternalServerError,
			"Message": fmt.Sprintf("Error creating person: %v", err),
		}
	} else {
		resp = map[string]any{
			"StatusCode": http.StatusOK,
			"Message": fmt.Sprintf("Person created with ID: %d", id),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var (
		resp map[string]any
		person model.Person
	)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Error converting ID to integer: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Printf("Error decoding person: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := repository.UpdatePerson(id, person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("More than one row affected: %d", rows)
		return
	}

	resp = map[string]any{
		"StatusCode": http.StatusOK,
		"Message": fmt.Sprintf("Person updated with ID: %d", id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	var resp map[string]any

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Error converting ID to integer: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := repository.DeletePerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("More than one row affected: %d", rows)
		return
	}

	resp = map[string]any{
		"StatusCode": http.StatusOK,
		"Message": fmt.Sprintf("Person deleted with ID: %d", id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func FindPersonByID(w http.ResponseWriter, r *http.Request) {
	var resp map[string]any

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error converting ID to integer: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	person, err := repository.FindPersonByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp = map[string]any{
		"StatusCode": http.StatusOK,
		"Person": person,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func FindAllPersons(w http.ResponseWriter, r *http.Request) {
	var resp map[string]any

	persons, err := repository.FindAllPersons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp = map[string]any{
		"StatusCode": http.StatusOK,
		"Persons": persons,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

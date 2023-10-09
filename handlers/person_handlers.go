package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"institute-person-api/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type PersonHandler struct {
	person models.PersonInterface
}

func NewPersonHandler(person models.PersonInterface) *PersonHandler {
	return &PersonHandler{person: person}
}

func (h *PersonHandler) AddPerson(responseWriter http.ResponseWriter, request *http.Request) {
	// transaction logging
	correltionId, _ := uuid.NewRandom()
	log.Printf("Begin CID: %s Add Person", correltionId)
	defer log.Printf("End CID: %s Add Person", correltionId)

	// Read the request body
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("ERROR CID: %s Bad Body Read %s", correltionId, err.Error())
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the new person document
	newPerson := h.person.PostPerson(body)

	// Return the new Person as JSON
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(newPerson)
}

func (h *PersonHandler) GetPerson(responseWriter http.ResponseWriter, request *http.Request) {
	// transaction logging
	correltionId, _ := uuid.NewRandom()
	log.Printf("Begin CID: %s Get Person", correltionId)
	defer log.Printf("End CID: %s Get Person", correltionId)

	// Get the Person ID from the path
	id := mux.Vars(request)["id"]

	// Get the Person from the database
	person := h.person.GetPerson(id)

	// Return the Person as JSON
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(person)
}

func (h *PersonHandler) GetPeople(responseWriter http.ResponseWriter, request *http.Request) {
	// transaction logging
	correltionId, _ := uuid.NewRandom()
	log.Printf("Begin CID: %s Get People", correltionId)
	defer log.Printf("End CID: %s Get People", correltionId)

	// Get all the people
	allPeople := h.person.GetAllNames()

	// Return the new Person as JSON
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(allPeople)
}

func (h *PersonHandler) UpdatePerson(responseWriter http.ResponseWriter, request *http.Request) {
	// transaction logging
	correltionId, _ := uuid.NewRandom()
	log.Printf("Begin CID: %s Update Person", correltionId)
	defer log.Printf("End CID: %s Update Person", correltionId)

	// Get the Person ID from the path
	id := mux.Vars(request)["id"]

	// Get the Request Body
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("ERROR CID: %s Bad Body Read %s", correltionId, err.Error())
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the person
	updatedPerson := h.person.PatchPerson(id, body)

	// Return the updated Person as JSON
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(updatedPerson)
}

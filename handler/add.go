package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/takumi616/go-english-vocabulary-api/entity"
	"github.com/takumi616/go-english-vocabulary-api/store"
)

type AddVocabulary struct {
	Store     *store.VocabularyStore
	Validator *validator.Validate
}

func (av *AddVocabulary) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	var requestBody struct {
		Title   string `json:"title" validate:"required"`
		Example string `json:"example" validate:"required"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		errRes := &ErrorResponse{Message: err.Error()}
		RespondJson(w, errRes, http.StatusInternalServerError)
		return
	}

	//Validate request body
	err = av.Validator.Struct(requestBody)
	if err != nil {
		errRes := &ErrorResponse{Message: err.Error()}
		RespondJson(w, errRes, http.StatusBadRequest)
		return
	}

	//Set request body to entity
	vocabulary := &entity.Vocabulary{
		Title:   requestBody.Title,
		Example: requestBody.Example,
		Created: time.Now(),
	}

	//Store new vocabulary (temp)
	vocabularyID, err := store.Vocabularies.AddVocabulary(vocabulary)
	if err != nil {
		errRes := &ErrorResponse{Message: err.Error()}
		RespondJson(w, errRes, http.StatusInternalServerError)
		return
	}

	//Generate response with returned vocabulary id
	response := struct {
		ID entity.VocabularyID `json:"id"`
	}{ID: vocabularyID}
	RespondJson(w, response, http.StatusCreated)
}

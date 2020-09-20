package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"go-library/api/library/models"
	"go-library/api/utils"
)

func (controller *LibraryController) GetAuthorByName(w http.ResponseWriter, r *http.Request) {
	// validations
	authorName, ok := mux.Vars(r)["author_name"]
	if !ok || len(strings.TrimSpace(authorName)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'author_name' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// gets author by name
	author, err := controller.service.GetAuthorByName(authorName)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&author)
	return
}

func (controller *LibraryController) GetBooksByAuthorName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// validations
	authorName, ok := mux.Vars(r)["author_name"]
	if !ok || len(strings.TrimSpace(authorName)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'author_name' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// gets books by author
	books, err := controller.service.GetBooksByAuthorName(authorName)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&books)
	return
}

func (controller *LibraryController) SaveNewAuthor(w http.ResponseWriter, r *http.Request) {
	newAuthor := models.Author{}
	rawErr := json.NewDecoder(r.Body).Decode(&newAuthor)
	if rawErr != nil {
		utils.HandleError(w, &utils.ApiError{
			Error: rawErr,
			Type:  http.StatusBadRequest,
		})
		return
	}

	// validations
	if len(newAuthor.Name) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("invalid author fields"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// sets author
	authorID, err := controller.service.SaveNewAuthor(newAuthor)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&authorID)
	return
}

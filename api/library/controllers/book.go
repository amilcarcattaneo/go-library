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

func (controller *LibraryController) GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	// validations
	title, ok := mux.Vars(r)["title"]
	if !ok || len(strings.TrimSpace(title)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'title' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// gets book by title
	book, err := controller.service.GetBookByTitle(title)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&book)
	return
}

type BookInput struct {
	models.Book
	AuthorID uint64 `json:"author_id"`
}

func (controller *LibraryController) SaveBook(w http.ResponseWriter, r *http.Request) {
	newBook := BookInput{}
	rawErr := json.NewDecoder(r.Body).Decode(&newBook)
	if rawErr != nil {
		utils.HandleError(w, &utils.ApiError{
			Error: rawErr,
			Type:  http.StatusBadRequest,
		})
		return
	}

	// validations
	if newBook.AuthorID == 0 || len(newBook.Title) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("invalid book fields"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// sets book
	bookID, err := controller.service.SaveNewBook(newBook.AuthorID, newBook.Book)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&bookID)
	return
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"go-library/api/utils"
)

func (controller *LibraryController) GetUserBooks(w http.ResponseWriter, r *http.Request) {
	// validations
	rawUserID, ok := mux.Vars(r)["user_id"]
	if !ok || len(strings.TrimSpace(rawUserID)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'user_id' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}
	userID, rawErr := strconv.Atoi(rawUserID)
	if rawErr != nil {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("invalid 'user_id' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// gets user books
	books, err := controller.service.GetUserAvailableBooks(uint64(userID))
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&books)
	return
}

func (controller *LibraryController) SaveNewUserBook(w http.ResponseWriter, r *http.Request) {
	// validations
	rawUserID, ok := mux.Vars(r)["user_id"]
	if !ok || len(strings.TrimSpace(rawUserID)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'user_id' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}
	userID, rawErr := strconv.Atoi(rawUserID)
	if rawErr != nil {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("invalid 'user_id' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	rawBookID, ok := mux.Vars(r)["book_id"]
	if !ok || len(strings.TrimSpace(rawBookID)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'book_id' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}
	bookID, rawErr := strconv.Atoi(rawBookID)
	if rawErr != nil {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("invalid 'book_id' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// save user book
	err := controller.service.SaveNewUserBook(uint64(userID), uint64(bookID))
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

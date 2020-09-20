package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"go-library/api/library/models"
	"go-library/api/utils"
)

type NewLoan struct {
	BookID       uint64 `json:"book_id"`
	UserIDLender uint64 `json:"user_id_lender"`
}

func (controller *LibraryController) RegisterLoan(w http.ResponseWriter, r *http.Request) {
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

	newLoan := NewLoan{}
	rawErr = json.NewDecoder(r.Body).Decode(&newLoan)
	if rawErr != nil {
		utils.HandleError(w, &utils.ApiError{
			Error: rawErr,
			Type:  http.StatusBadRequest,
		})
		return
	}
	if newLoan.BookID == 0 || newLoan.UserIDLender == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("invalid payload"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// check user has book available
	books, err := controller.service.GetUserAvailableBooks(newLoan.UserIDLender)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	if !userHasBookAvailable(books, newLoan.BookID) {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("book not available"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	// saves loan and updates user book not available
	err = controller.service.SaveLoan(uint64(userID), newLoan.BookID, newLoan.UserIDLender)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func userHasBookAvailable(books []models.Book, bookID uint64) bool {
	for _, book := range books {
		if book.ID == bookID {
			return true
		}
	}
	return false
}

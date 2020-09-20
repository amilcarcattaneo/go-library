package services

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"

	"go-library/api/library/models"
	"go-library/api/utils"
)

const (
	getUserBooks    = `SELECT ub.book_id FROM user_books AS ub WHERE ub.user_id = ? AND ub.available = TRUE`
	saveNewUserBook = `INSERT INTO user_books (user_id, book_id) VALUES (?, ?)`
)

var (
	// ErrUserNotFound error returned in case user is not found
	ErrUserNotFound = errors.New("user not found")
)

type bookID struct {
	BookID uint64 `json:"book_id"`
}

func (service *libraryService) GetUserAvailableBooks(id uint64) ([]models.Book, *utils.ApiError) {
	bookIDs := []bookID{}
	err := service.db.Raw(getUserBooks, id).Scan(&bookIDs).Error
	if err != nil {
		rawtype := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			rawtype = http.StatusNotFound
		}
		return nil, &utils.ApiError{
			Error: err,
			Type:  rawtype,
		}
	}
	books := []models.Book{}
	for _, bookID := range bookIDs {
		book, err := service.GetBookByID(bookID.BookID)
		if err != nil {
			return nil, err
		}
		books = append(books, *book)
	}

	return books, nil
}

func (service *libraryService) UpdateUserBookAvailability(tx *gorm.DB, userID, bookID uint64, availability bool) *utils.ApiError {
	db := service.db
	if tx != nil {
		db = tx
	}

	if err := db.Model(&models.UserBook{
		UserID: userID,
		BookID: bookID,
	}).Select("available").Updates(map[string]interface{}{"available": availability}).Error; err != nil {
		return &utils.ApiError{
			Error: err,
			Type:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (service *libraryService) SaveNewUserBook(userID, bookID uint64) *utils.ApiError {
	if err := service.db.Exec(saveNewUserBook, userID, bookID).Error; err != nil {
		rawtype := http.StatusInternalServerError
		if strings.Contains(err.Error(), "user_id") {
			err = ErrUserNotFound
			rawtype = http.StatusNotFound
		}
		if strings.Contains(err.Error(), "book_id") {
			err = ErrBookNotFound
			rawtype = http.StatusNotFound
		}
		return &utils.ApiError{
			Error: err,
			Type:  rawtype,
		}
	}

	return nil
}

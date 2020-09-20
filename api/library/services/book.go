package services

import (
	"fmt"
	"net/http"
	"strings"

	"go-library/api/library/models"
	"go-library/api/utils"

	"github.com/pkg/errors"
)

const (
	getBooksByAuthorID = `SELECT b.* FROM books AS b INNER JOIN author_books AS ab ON ab.book_id = b.id WHERE ab.author_id = ?`
	saveAuthorBook     = `INSERT INTO author_books (author_id, book_id) VALUES (?, ?)`
	getBookAuthor      = `SELECT a.* FROM authors AS a INNER JOIN author_books AS ab ON ab.author_id = a.id WHERE ab.book_id = ?`
)

var (
	// ErrBookNotFound error returned in case the book is not found
	ErrBookNotFound = errors.New("book not found")
	// ErrBookNotSaved error returned in case the book couldn't be saved
	ErrBookNotSaved = errors.New("book not saved")
)

func (service *libraryService) GetBookByTitle(title string) ([]models.Book, *utils.ApiError) {
	books := []models.Book{}

	if err := service.db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", getString(title))).Find(&books).Error; err != nil {
		rawtype := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			rawtype = http.StatusNotFound
			err = ErrBookNotFound
		}
		return nil, &utils.ApiError{
			Error: err,
			Type:  rawtype,
		}
	}

	for idx, book := range books {
		if err := service.db.Raw(getBookAuthor, book.ID).Scan(&books[idx].Author).Error; err != nil {
			rawtype := http.StatusInternalServerError
			if strings.Contains(err.Error(), "not found") {
				rawtype = http.StatusNotFound
				err = ErrAuthorNotFound
			}
			return nil, &utils.ApiError{
				Error: err,
				Type:  rawtype,
			}
		}
	}

	return books, nil
}

func (service *libraryService) GetBookByID(id uint64) (*models.Book, *utils.ApiError) {
	var book models.Book
	if err := service.db.Where(&models.Book{ID: id}).Find(&book).Error; err != nil {
		rawtype := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			rawtype = http.StatusNotFound
			err = ErrBookNotFound
		}
		return nil, &utils.ApiError{
			Error: err,
			Type:  rawtype,
		}
	}

	if err := service.db.Raw(getBookAuthor, book.ID).Scan(&book.Author).Error; err != nil {
		rawtype := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			rawtype = http.StatusNotFound
			err = ErrAuthorNotFound
		}
		return nil, &utils.ApiError{
			Error: err,
			Type:  rawtype,
		}
	}

	return &book, nil
}

func (service *libraryService) GetBooksByAuthorName(authorName string) ([]models.Book, *utils.ApiError) {
	author, err := service.GetAuthorByName(authorName)
	if err != nil {
		return nil, err
	}
	// TODO: Revisar
	return service.GetBooksByAuthorID(author[0].ID)
}

func (service *libraryService) GetBooksByAuthorID(id uint64) ([]models.Book, *utils.ApiError) {
	books := []models.Book{}
	err := service.db.Raw(getBooksByAuthorID, id).Scan(&books).Error
	if err != nil {
		rawtype := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			rawtype = http.StatusNotFound
			err = ErrBookNotFound
		}
		return nil, &utils.ApiError{
			Error: err,
			Type:  rawtype,
		}
	}
	author, _ := service.GetAuthorByID(id)
	for idx := range books {
		books[idx].Author = *author
	}

	return books, nil
}

func (service *libraryService) SaveNewBook(authorID uint64, book models.Book) (uint64, *utils.ApiError) {
	tx := service.db.Begin()
	if err := tx.Create(&book).Error; err != nil {
		tx.Rollback()
		return 0, &utils.ApiError{
			Error: ErrBookNotSaved,
			Type:  http.StatusInternalServerError,
		}
	}

	if err := tx.Exec(saveAuthorBook, authorID, book.ID).Error; err != nil {
		tx.Rollback()
		rawtype := http.StatusInternalServerError
		if strings.Contains(err.Error(), "author_id") {
			err = ErrAuthorNotFound
			rawtype = http.StatusNotFound
		}
		return 0, &utils.ApiError{
			Error: err,
			Type:  rawtype,
		}
	}

	tx.Commit()
	return book.ID, nil
}

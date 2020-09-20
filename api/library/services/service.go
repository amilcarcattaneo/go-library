package services

import (
	"regexp"

	"github.com/jinzhu/gorm"

	"go-library/api/library/models"
	"go-library/api/utils"
)

type LibraryService interface {
	GetBookByTitle(title string) ([]models.Book, *utils.ApiError)
	GetBookByID(id uint64) (*models.Book, *utils.ApiError)
	SaveNewBook(authorID uint64, book models.Book) (uint64, *utils.ApiError)

	GetAuthorByName(authorName string) ([]models.Author, *utils.ApiError)
	GetAuthorByID(id uint64) (*models.Author, *utils.ApiError)
	SaveNewAuthor(author models.Author) (uint64, *utils.ApiError)

	GetBooksByAuthorName(authorName string) ([]models.Book, *utils.ApiError)
	GetBooksByAuthorID(id uint64) ([]models.Book, *utils.ApiError)

	GetUserAvailableBooks(id uint64) ([]models.Book, *utils.ApiError)
	UpdateUserBookAvailability(tx *gorm.DB, userID, bookID uint64, availability bool) *utils.ApiError
	SaveNewUserBook(userID, bookID uint64) *utils.ApiError

	SaveLoan(userID, bookID, userIDLender uint64) *utils.ApiError
}

type libraryService struct {
	db *gorm.DB
}

func NewLibraryService(db *gorm.DB) LibraryService {
	return &libraryService{
		db: db,
	}
}

var re = regexp.MustCompile(`\W+`)

func getString(s string) string {
	return re.ReplaceAllString(s, "%")
}

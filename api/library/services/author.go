package services

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"go-library/api/library/models"
	"go-library/api/utils"
)

var (
	// ErrAuthorNotFound error returned in case the author is not found
	ErrAuthorNotFound = errors.New("author not found")
	// ErrAuthorNotSaved error returned in case the author couldn't be saved
	ErrAuthorNotSaved = errors.New("author not saved")
)

func (service *libraryService) GetAuthorByName(authorName string) ([]models.Author, *utils.ApiError) {
	authors := []models.Author{}

	if err := service.db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", getString(authorName))).Find(&authors).Error; err != nil {
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
	return authors, nil
}

func (service *libraryService) GetAuthorByID(id uint64) (*models.Author, *utils.ApiError) {
	var author models.Author
	if err := service.db.Where(&models.Author{ID: id}).Find(&author).Error; err != nil {
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
	return &author, nil
}

func (service *libraryService) SaveNewAuthor(author models.Author) (uint64, *utils.ApiError) {
	if err := service.db.Create(&author).Error; err != nil {
		return 0, &utils.ApiError{
			Error: ErrAuthorNotSaved,
			Type:  http.StatusInternalServerError,
		}
	}

	return author.ID, nil
}

package services

import (
	"net/http"
	"time"

	"go-library/api/library/models"
	"go-library/api/utils"
)

func (service *libraryService) SaveLoan(userID, bookID, userIDLender uint64) *utils.ApiError {
	userLoan := &models.UserLoan{
		UserID:       userID,
		BookID:       bookID,
		DueDate:      time.Now().AddDate(0, 1, 0),
		UserIDLender: userIDLender,
	}

	tx := service.db.Begin()
	if rawErr := tx.Create(userLoan).Error; rawErr != nil {
		tx.Rollback()
		return &utils.ApiError{
			Error: rawErr,
			Type:  http.StatusInternalServerError,
		}
	}

	if err := service.UpdateUserBookAvailability(tx, userIDLender, bookID, false); err != nil {
		tx.Rollback()
		return err
	}
	if rawErr := tx.Commit().Error; rawErr != nil {
		return &utils.ApiError{
			Error: rawErr,
			Type:  http.StatusInternalServerError,
		}
	}

	return nil
}

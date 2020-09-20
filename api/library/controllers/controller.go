package controllers

import (
	"github.com/jinzhu/gorm"

	"go-library/api/library/services"
)

type LibraryController struct {
	service services.LibraryService
}

func NewLibraryController(db *gorm.DB) *LibraryController {
	return &LibraryController{
		service: services.NewLibraryService(db),
	}
}

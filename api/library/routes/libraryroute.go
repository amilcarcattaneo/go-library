package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"go-library/api/library/controllers"
)

func Init(router *mux.Router, db *gorm.DB) {
	libraryController := controllers.NewLibraryController(db)

	router.HandleFunc("/books/{title}", libraryController.GetBookByTitle).Methods(http.MethodGet)
	router.HandleFunc("/books", libraryController.SaveBook).Methods(http.MethodPost)

	router.HandleFunc("/authors/{author_name}", libraryController.GetAuthorByName).Methods(http.MethodGet)
	router.HandleFunc("/authors/{author_name}/books", libraryController.GetBooksByAuthorName).Methods(http.MethodGet)
	router.HandleFunc("/authors", libraryController.SaveNewAuthor).Methods(http.MethodPost)

	router.HandleFunc("/users/{user_id}/loan", libraryController.RegisterLoan).Methods(http.MethodPut)
	router.HandleFunc("/users/{user_id}/books/{book_id}", libraryController.SaveNewUserBook).Methods(http.MethodPut)
	router.HandleFunc("/users/{user_id}/books", libraryController.GetUserBooks).Methods(http.MethodGet)
}

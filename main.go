package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"

	"go-library/api/library/routes"
)

func main() {
	routerHandler := mux.NewRouter()
	routerHandler.Use(commonMiddleware)

	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", os.Getenv("DBUsername"),
			os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBLibraryName")),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	routes.Init(routerHandler, db)

	handler := cors.Default().Handler(routerHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

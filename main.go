package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	h "github.com/parikshitg/goauth2/handlers"
	"github.com/parikshitg/goauth2/models"
)

func main() {

	var err error
	models.Db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test sslmode=disable")
	if err != nil {

		log.Println("failed to connect database:  err : ", err)
		return
	}

	defer models.Db.Close()

	log.Println("DB Connected Successfully")

	// Create Table
	models.Db.AutoMigrate(&models.User{})

	// Routes
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/github/login", h.GithubLogin)
	http.HandleFunc("/github/callback", h.GithubCallback)
	http.HandleFunc("/dashboard", h.Dashboard)
	http.HandleFunc("/register", h.Register)

	// Serving static files
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("static/css"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("static/js"))))

	http.ListenAndServe(":8080", nil)
}

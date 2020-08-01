package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	h "github.com/parikshitg/goauth2/handlers"
	"github.com/parikshitg/goauth2/models"
	"github.com/parikshitg/goauth2/sessions"
)

func main() {

	r := mux.NewRouter()

	var err error
	models.Db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test sslmode=disable")
	if err != nil {

		log.Println("failed to connect database:  err : ", err)
		return
	}

	defer models.Db.Close()

	log.Println("DB Connected Successfully")

	models.Db.Debug().DropTableIfExists(&models.User{})
	// Create Table
	models.Db.AutoMigrate(&models.User{})

	// Routes
	r.HandleFunc("/", h.Home)
	r.HandleFunc("/login", sessions.UnauthenticatedUser(h.Login))
	r.HandleFunc("/github/login", h.GithubLogin)
	r.HandleFunc("/github/callback", h.GithubCallback)
	r.HandleFunc("/linkedin/login", h.LinkedinLogin)
	r.HandleFunc("/linkedin/callback", h.LinkedinCallback)
	r.HandleFunc("/user/all", sessions.AuthenticatedUser(h.Dashboard))
	r.HandleFunc("/user/{id:[0-9]+}", sessions.AuthenticatedUser(h.UserDetailsHandler))
	r.HandleFunc("/user/search", sessions.AuthenticatedUser(h.Search))
	r.HandleFunc("/user/set_password", sessions.AuthenticatedUser(h.SetPasswordHandler))
	r.HandleFunc("/register", sessions.UnauthenticatedUser(h.Register))
	r.HandleFunc("/logout", sessions.AuthenticatedUser(h.Logout))

	// Serving static files
	fcss := http.FileServer(http.Dir("static/css"))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", fcss))
	fjs := http.FileServer(http.Dir("static/js"))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", fjs))

	http.ListenAndServe(":8080", r)
}

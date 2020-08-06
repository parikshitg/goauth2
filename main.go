package main

import (
	"log"
	"net/http"

	"github.com/dghubble/gologin/twitter"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"goauth2/conf"
	h "goauth2/handlers"
	"goauth2/models"
	"goauth2/sessions"
)

func main() {

	r := mux.NewRouter()

	var err error
	models.Db, err = gorm.Open(conf.Db, "host="+conf.Dbhost+" port="+conf.Dbport+" user="+conf.Dbuser+" dbname="+conf.Dbname+" sslmode=disable")
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
	r.Handle("/twitter/login", twitter.LoginHandler(h.Oauth1Config, nil))
	r.Handle("/twitter/callback", twitter.CallbackHandler(h.Oauth1Config, h.IssueSession(), nil))
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

package main

import (
	"net/http"

	h "github.com/parikshitg/goauth2/handlers"
)

func main() {

	// Routes
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/github/login",h.GithubLogin)
	http.HandleFunc("/github/callback", h.GithubCallback)
	http.HandleFunc("/dashboard", h.Dashboard)
	http.HandleFunc("/register", h.Register)

	// Serving static files
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("static/css"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("static/js"))))

	http.ListenAndServe(":8080", nil)
}

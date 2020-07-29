package main

import (
	"net/http"

	h "github.com/parikshitg/goauth2/handlers"
)

func main() {

	http.HandleFunc("/", h.Home)
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/dashboard", h.Dashboard)

	http.ListenAndServe(":8080", nil)
}

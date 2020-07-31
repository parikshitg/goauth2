package sessions

import (
	"net/http"

	s "github.com/gorilla/sessions"
)

// initialize session key
var key = []byte("super-secret-key")
var Store = s.NewCookieStore(key)

// Checks if user is authenticated (middleware)
func AuthenticatedUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := Store.Get(r, "auth-cookie")

		cookie, ok := session.Values["authenticated"]

		if cookie == true && ok {
			f(w, r)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	}
}

// checks unauthenticated user (middleware)
func UnauthenticatedUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := Store.Get(r, "auth-cookie")

		cookie, ok := session.Values["authenticated"]

		if cookie == true && ok {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}

		f(w, r)
	}
}

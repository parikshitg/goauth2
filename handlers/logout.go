package handlers

import (
	"net/http"

	s "goauth2/sessions"
)

// Logout function
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Store.Get(r, "auth-cookie")

	session.Values["Useremail"] = ""
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

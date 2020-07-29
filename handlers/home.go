package handlers

import (
	"net/http"
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

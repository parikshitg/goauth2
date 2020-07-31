package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/parikshitg/goauth2/models"
	s "github.com/parikshitg/goauth2/sessions"
)

// Dashboard Handler
func Dashboard(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["title"] = "Dashboard"

	page, err := template.ParseFiles("templates/dashboard.html", "templates/footer.html", "templates/header2.html")
	if err != nil {
		log.Println("ParseFiles: ", err)
		return
	}

	session, _ := s.Store.Get(r, "auth-cookie")
	username, _ := session.Values["User"]
	data["User"] = username

	// Display List of Users
	userList := models.UsersTable()
	data["Users"] = userList

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

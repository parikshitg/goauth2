package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/parikshitg/goauth2/models"
	s "github.com/parikshitg/goauth2/sessions"
)

// SetPassword Handler
func SetPassword(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["title"] = "Set Password"

	page, err := template.ParseFiles("templates/set-password.html", "templates/footer.html", "templates/header2.html")
	if err != nil {
		log.Println("ParseFiles: ", err)
		return
	}

	session, _ := s.Store.Get(r, "auth-cookie")
	useremail, _ := session.Values["Useremail"]
	user, _ := models.ExistingUser(useremail.(string))
	data["User"] = user.Name

	if r.Method == http.MethodPost {

		// password := r.FormValue("password")
		// password2 := r.FormValue("password2")
		// password3 := r.FormValue("password3")

	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

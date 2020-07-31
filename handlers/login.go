package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/parikshitg/goauth2/models"
	s "github.com/parikshitg/goauth2/sessions"
)

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["title"] = "Home"

	page, err := template.ParseFiles("templates/login.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		log.Println("ParseFiles: ", err)
		return
	}

	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		msg, ok := LoginUser(email, password)
		if !ok {

			data["Flash"] = msg
			err = page.Execute(w, data)
			if err != nil {
				log.Fatal("Execute:", err)
			}
			return
		}

		user, _ := models.ExistingUser(email)

		// setting up a session
		session, err := s.Store.Get(r, "auth-cookie")
		if err != nil {
			log.Println("Session Error:", err)
			return
		}
		session.Values["User"] = user.Name
		session.Save(r, w)

		http.Redirect(w, r, "/user/all", http.StatusSeeOther)
		return
	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

// Logins the User
func LoginUser(email, password string) (string, bool) {

	if email == "" || password == "" {

		flash.Message = "Fields Can not be empty !!"
		return flash.Message, false
	}

	user, exists := models.ExistingUser(email)
	if !exists {

		flash.Message = "Invalid Email !!"
		return flash.Message, false
	}

	if password != user.Password {
		flash.Message = "Invalid User !!"
		return flash.Message, false
	}

	log.Println("Logged in Successfully")
	return flash.Message, true
}

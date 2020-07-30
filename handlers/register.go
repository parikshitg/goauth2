package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/parikshitg/goauth2/models"
)

// Register handler
func Register(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["title"] = "Register"

	page, err := template.ParseFiles("templates/register.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		log.Println("ParseFiles: ", err)
		return
	}

	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		password2 := r.FormValue("password2")
		log.Println("name : ", name, "email : ", email, "password : ", password, "password2 : ", password2)

		ok := models.ExistingUser(email)
		if ok {
			log.Println("User Existing already")
		} else {
			models.CreateUser(name, email, password)
		}
	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

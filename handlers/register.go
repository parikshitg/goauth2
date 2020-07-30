package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/parikshitg/goauth2/models"
)

type Flash struct {
	Message string
}

var flash Flash

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

		// Register the User
		res := RegisterUser(name, email, password, password2)
		data["Flash"] = res
		log.Println("flash message : ", res)
	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

// Registers User
func RegisterUser(name, email, password, password2 string) string {

	if name == "" || email == "" || password == "" || password2 == "" {

		flash.Message = "Fields Can not be empty !!"
		return flash.Message
	}

	ok := models.ExistingUser(email)
	if ok {
		flash.Message = "User Already Registered !!"
		return flash.Message
	}

	if password != password2 {
		flash.Message = "Passwords Doesn't Match !!"
		return flash.Message
	}

	models.CreateUser(name, email, password)
	flash.Message = "Registered Successfully."

	return flash.Message
}

package handlers

import (
	"html/template"
	"log"
	"net/http"
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

		username := r.FormValue("username")
		password := r.FormValue("password")
		password2 := r.FormValue("password2")
		log.Println("username : ", username, "password : ", password,"password2 : ", password2)
	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

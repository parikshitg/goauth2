package handlers

import (
	"html/template"
	"log"
	"net/http"
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

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
